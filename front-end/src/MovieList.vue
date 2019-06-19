<template>
  <div class="content">
    <div id="movies" class="movies">
      <h2>Welcome to Techveritos Multiplex</h2>
      <div v-if="show" class="row">
        <div v-for="(movie,index) in movies" class="card" @click="calling(index+1)">
          <img :src="getMoviePosterUrl(movie.url)" :alt="movie.url" class="full-width">
          <h2>{{movie.movie_name}}</h2>
          <h4>{{movie.description}}</h4>
          <p>Show{{index+1}}</p>
          <p class="card-button"><button>Book Now</button></p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Booking from './Booking.vue';
import axios from 'axios'
import VueAxios from 'vue-axios'
export default {
  components: {
    booking: Booking,
  },
  data: function() {
    return {
      movieID: 0,
      movies: [],
      link: '',
      show: true,
    };
  },
  created: function() {
    var url = "http://localhost:8000/api/getMovieList";
    axios
      .get(url)
      .then(response => {
        this.movies = response.data;
      })
  },
  methods: {
    calling: function(id) {
      this.movieID = id;
      this.link = '/booking/' + this.movieID;
      this.show = false;
      this.$router.push(this.link);
    },
    getMoviePosterUrl: function(url) {
      return "/src/assets/" + url
    }
  }
}

</script>
