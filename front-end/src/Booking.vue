<template>
  <div class="content">
    <div id="shows" class="col-md-10">
      <div class="container">
        <h2>
            <a class="black pointer" @click="goHome()"><i class="fas fa-arrow-left arrows"></i>
            Show {{showID}}</a>
            </h2>
        <div class="container">
          <table class="showsTable">
            <thead>
              <tr>
                <th colspan="9"><h3 class="conatiner-small">Chose your seats for show {{showID}}</h3>
                </th>
              </tr>
            </thead>
            <tbody v-for="item in showDetails.type" class="">
              <tr>
                <td colspan="4" class="table-decoration">
                  <p v-if="item.name=='A'">Rs {{item.price}}-Platinum Seats</p>
                  <p v-if="item.name=='B'">Rs {{item.price}}-Gold Seats</p>
                  <p v-if="item.name=='C'">Rs {{item.price}}-Silver Seats</p>
                </td>
              </tr>
              <tr>
                <td v-for="(sets,index) in item.seat" v-if="sets.Active" class="box" :id="sets.seat_no" @click="seatClicked(item,sets)" v-bind:class="{ ticketTaken: sets.taken , ticketNotTaken: !sets.taken}">
                  <p class="column pointer center margin">
                    {{sets.seat_no}}
                  </p>
                </td>
                <td v-else>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="container">
          <h2 class="pricing">Pricing</h2>
          <table class="pricingTable">
            <tr>
              <td>Ticket Price</td>
              <td>{{ticket.price}}</td>
            </tr>
            <tr>
              <td>Service Tax(14% on ticket price)</td>
              <td>{{ticket.st}}</td>
            </tr>
            <tr>
              <td>Swachh Bharat Cess(0.5% on Service Tax)</td>
              <td>{{ticket.sbc}}</td>
            </tr>
            <tr>
              <td>Krishi Kalyan Cess(0.5% on Service Tax)</td>
              <td>{{ticket.kkc}}</td>
            </tr>
            <tr>
              <td><b>Final Amount</b></td>
              <td><b>{{ticket.final}}</b></td>
            </tr>
          </table>
          <div class="center conatiner-small">
            <button @click="bookSeats" class="button"><b>BOOK SEATS</b></button>
            <h5 class="conatiner-small green">{{message}}</h5>
          </div>
        </div>
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
      showID: 0,
      showDetails: '',
      difference: 0,
      ticket: {
        price: 0,
        refresh: function() {
          this.price = 0;
          this.calculateTax();
        },
        init: function(price) {
          this.price = price;
          this.calculateTax();
        },
        getPrice: function() {
          return this.price;
        },
        calculateTax: function() {
          this.st = parseFloat(((this.price * 14) / 100).toFixed(2));
          this.sbc = parseFloat(((this.st * 0.5) / 100).toFixed(2));
          this.kkc = parseFloat(((this.st * 0.5) / 100).toFixed(2));
          this.final = parseFloat(((this.price) + (this.st) + (this.sbc) + (this.kkc)).toFixed(2));
        }
      },
      seatsSelected: [],
      message: '',
    };
  },
  created: function() {
    this.showID = this.$route.params.showID;
    this.refreshSeats();
    this.ticket.init(0);
  },
  methods: {
    goHome: function() {
      this.$router.push('/');
    },
    refreshSeats: function() {
      axios
        .get("http://localhost:8000/api/getMultiplex")
        .then(response => {
          this.showDetails = JSON.parse(JSON.stringify(response.data[this.showID - 1]));
        });
    },
    seatClicked: function(set, seat) {
      if (!seat.taken) {
        if (this.seatsSelected.includes(seat.seat_no)) {
          var index = this.seatsSelected.indexOf(seat.seat_no);
          if (index > -1) {
            this.seatsSelected.splice(index, 1);
            this.ticket.init(this.ticket.getPrice() - set.price);
            document.getElementById(seat.seat_no).style.setProperty("background-color", "white", "important");
          }
        } else {
          this.seatsSelected.push(seat.seat_no);
          this.ticket.init(this.ticket.getPrice() + set.price);
          document.getElementById(seat.seat_no).style.setProperty("background-color", "green");
        }
      }
    },
    calculateTax: function(price) {
      var result = new Object();
      result.price = price;
      result.st = parseFloat(((result.price * 14) / 100).toFixed(2));
      result.sbc = parseFloat(((result.st * 0.5) / 100).toFixed(2));
      result.kkc = parseFloat(((result.st * 0.5) / 100).toFixed(2));
      result.final = parseFloat(((result.price) + (result.st) + (result.sbc) + (result.kkc)).toFixed(2));
      return result;
    },
    bookSeats: function() {
      if (this.seatsSelected.length > 0) {
        var url = "http://localhost:8000/api/bookShow/" + this.showID + "/" + this.seatsSelected;
        axios
          .get(url)
          .then(response => {
            if (response.data.success) {
              this.message = "Successfully booked " + this.seatsSelected + ". Enjoy the movie!";
              alert("Successfully Booked " + this.seatsSelected.length + " seats for you");
              this.seatsSelected = [];
              this.ticket.refresh();
              this.refreshSeats();
            }
          })
          .catch(error => {
            console.log(error);
          });
      } else {
        alert("Select some seats first");
      }
    }
  }
}

</script>
