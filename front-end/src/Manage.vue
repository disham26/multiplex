<template>
  <div id="shows" class="content">
    <div class="container">
      <h2>Manage Section</h2>
      <div class="container ">
        <h2 class="pricing">Overall</h2>
        <table class="pricingTable managePricingTable">
          <tr>
            <td>Revenue</td>
            <td>{{overall.price}}</td>
          </tr>
          <tr>
            <td>Service Tax(14% on ticket price)</td>
            <td>{{overall.st}}</td>
          </tr>
          <tr>
            <td>Swachh Bharat Cess(0.5% on Service Tax)</td>
            <td>{{overall.sbc }}</td>
          </tr>
          <tr>
            <td>Krishi Kalyan Cess(0.5% on Service Tax)</td>
            <td>{{overall.kkc}}</td>
          </tr>
          <tr>
            <td><b>Final Amount</b></td>
            <td><b>{{overall.final}}</b></td>
          </tr>
        </table>
      </div>
      <div v-for="(item,index) in shows" class="container">
        <h2 class="pricing">Show {{index+1}}</h2>
        <table class="pricingTable managePricingTable">
          <tr>
            <td>Revenue</td>
            <td>{{item.price}}</td>
          </tr>
          <tr>
            <td>Service Tax(14% on ticket price)</td>
            <td>{{item.st}}</td>
          </tr>
          <tr>
            <td>Swachh Bharat Cess(0.5% on Service Tax)</td>
            <td>{{item.sbc}}</td>
          </tr>
          <tr>
            <td>Krishi Kalyan Cess(0.5% on Service Tax)</td>
            <td>{{item.kkc}}</td>
          </tr>
          <tr>
            <td><b>Final Amount</b></td>
            <td><b>{{item.final}}</b></td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
import VueAxios from 'vue-axios'
export default {
  data: function() {
    return {
      showDetails: [],
      overall: {},
      shows: [],
      price: {
        init: function(price) {
          this.price = price;
          this.calculateTax();
        },
        calculateTax: function() {
          this.st = parseFloat(((this.price * 14) / 100).toFixed(2));
          this.sbc = parseFloat(((this.st * 0.5) / 100).toFixed(2));
          this.kkc = parseFloat(((this.st * 0.5) / 100).toFixed(2));
          this.final = parseFloat(((this.price) + (this.st) + (this.sbc) + (this.kkc)).toFixed(2));
        }
      }
    };
  },
  created: function() {
    this.getRevenue();
  },
  methods: {
    getRevenue: function() {
      var url = "http://localhost:8000/api/getShows";
      axios
        .get(url)
        .then(response => {
          this.showDetails = response.data;
          this.updateTotal(this.showDetails);
        })
    },
    updateTotal: function(data) {
      var totalPrice = 0;
      var temp;
      var tempPrice = 0;
      for (let i = 0; i < data.length; i++) {
        for (let j = 0; j < data[i].type.length; j++) {
          for (let k = 0; k < data[i].type[j].seat.length; k++) {
            if (data[i].type[j].seat[k].taken) {
              tempPrice += data[i].type[j].price;
              totalPrice += data[i].type[j].price;
            }
          }
        }
        temp = Object.create(this.price);
        temp.init(tempPrice);
        this.shows.push(temp);
        tempPrice = 0;
      }
      this.overall = Object.create(this.price);
      this.overall.init(totalPrice);
    },
  },
}

</script>
