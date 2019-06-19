import Vue from 'vue';
import VueRouter from 'vue-router';
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './App.vue';
import Booking from './Booking.vue';
import Manage from './Manage.vue';
import MovieList from './MovieList.vue';
import NotFound from './NotFound.vue';

Vue.use(VueRouter, VueAxios, axios);
const routes = [
  { path: '/booking/:showID', component: Booking },
  { path: '/manage', component: Manage },
  { path: '/', component: MovieList },
  { path: "*", component: NotFound }
];
const router = new VueRouter({
  routes,
});
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
