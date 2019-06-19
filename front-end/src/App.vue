<template>
  <div>
    <appHeader></appHeader>
    <router-view></router-view>
  </div>
</template>
<script>
import Header from './shared/Header.vue';
import Manage from './Manage.vue';
import MovieList from './MovieList.vue';
export default {
  data: function() {
    return {
      home: true,
    };
  },
  created: function() {
    this.$router.push("/");
  },
  methods: {
    update: function() {

    }
  },
  components: {
    appHeader: Header,
    manage: Manage,
    movieList: MovieList,
  },
}

</script>
