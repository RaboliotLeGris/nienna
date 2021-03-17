<template>
  <div id="session">
    <span v-if="this.$store.state.logged">
      <button v-on:click="logout"> disconnect </button>
      <span> {{ this.$store.state.username }} </span>
    </span>
    <span v-else>
      <router-link to="/register">Register</router-link> |
      <router-link to="/login">Login</router-link>
    </span>
  </div>
</template>

<script>
import cookies from 'js-cookie';

export default {
  name: 'Session',
  methods: {
    logout() {
      cookies.remove('nienna_username');
      cookies.remove('nienna');
      this.$store.commit('logout');
    },
  },
};
</script>

<style>
  #session {
    float: right;
  }
</style>
