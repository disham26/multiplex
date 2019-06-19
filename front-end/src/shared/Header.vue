<template>
  <div>
    <div class="topnav" id="myTopnav">
      <section @click="toggleHome(true)">
        <router-link v-bind:class="{ active: home }" to="/">Home</router-link>
      </section>
      <section @click="toggleHome(false)">
        <router-link v-bind:class="{ active: !home }" to="/manage">Manage</router-link>
      </section>
      <section href="javascript:void(0);" class="icon" @click="callBar()">
        <i class="fa fa-bars"></i>
      </section>
    </div>
  </div>
</template>
<script>
export default {
  data: function() {
    return {
      home: true,
    }
  },
  created: function() {
    this.home = true;
  },
  methods: {
    callBar: function() {
      var x = document.getElementById("myTopnav");
      if (x.className === "topnav") {
        x.className += " responsive";
      } else {
        x.className = "topnav";
      }
    },
    toggleHome: function(result) {
      this.home = result;
    },
  }
}

</script>
<style>
body {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
}

.movies {
  margin-top: 50px;
}

.topnav {
  overflow: hidden;
  background-color: #333;
}

.row {
  display: -ms-flexbox;
  /* IE10 */
  display: flex;
  -ms-flex-wrap: wrap;
  /* IE10 */
  flex-wrap: wrap;
  padding: 0 4px;
}

.card-button {
  margin: 0;
}

.topnav a {
  float: left;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.center {
  text-align: center;
}

.margin {
  margin: 0 10px;
}

.topnav a:hover {
  background-color: #ddd;
  color: black;
}

.full-width {
  width: 100%;
}

.container {
  margin-top: 50px;
}

.culmn {
  position: relative;
}

.container-small {
  margin-top: 20px;
}

.pointer {
  cursor: pointer;
}

.active {
  background-color: #4CAF50;
  color: white;
}

.topnav .icon {
  display: none;
}

.content {
  margin: 0 50px;
}

.table-decoration {
  text-decoration: underline;
}

@media screen and (max-width: 600px) {
  .topnav a:not(:first-child) {
    display: none;
  }

  .content {
    margin: 0 20px;
  }

  .topnav a.icon {
    float: right;
    display: block;
  }

  .topnav.responsive {
    position: relative;
  }

  .topnav.responsive .icon {
    position: absolute;
    right: 0;
    top: 0;
  }

  .topnav.responsive a {
    float: none;
    display: block;
    text-align: left;
  }

  .card {
    -ms-flex: 100%;
    flex: 100%;
    max-width: 100%;
  }
}

.showsTable {
  margin: auto;
  width: 50%;

}

.arrows {
  color: black;
}

.black {
  color: black;
}

.pricingTable {
  margin: auto;
  width: 70%;

}

.managePricingTable {
  border-bottom: 1px solid #ddd;
}

.pricing {
  margin: auto;
  width: 70%;
  margin-bottom: 50px;
}

.green {
  color: green;
}

.button {
  background-color: #4CAF50;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  width: 300px;
}

.ticketTaken {
  background-color: gray !important;
  cursor: pointer;
}

.ticketNotTaken {
  background-color: white;
  color: black;
}

.box {
  width: auto;
  border: 1px solid;
  border-spacing: 10px;
  border-collapse: separate;
}

.card {
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  max-width: 300px;
  text-align: center;
  font-family: arial;
  margin: 20px auto;
}

.manage-card {
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);

  text-align: center;
  font-family: arial;
  margin: 20px auto;
}

.card button {
  border: none;
  outline: 0;
  padding: 12px;
  color: white;
  background-color: #000;
  text-align: center;
  cursor: pointer;
  width: 100%;
  font-size: 18px;
}

.card button:hover {
  opacity: 0.7;
}

</style>
