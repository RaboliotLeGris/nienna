<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link>
      <router-link v-show="this.$store.state.logged" to="/upload">Upload</router-link>
      <Session/>
    </div>
    <router-view/>
  </div>
</template>

<script>
import Session from '@/components/Session.vue';

export default {
  name: 'App',
  components: {
    Session,
  },
  created() {
    document.title = 'Nienna - Watch Videos';
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
