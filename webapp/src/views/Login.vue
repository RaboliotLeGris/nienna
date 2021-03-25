<template>
  <div class="login">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <p>
        <label for="username">Username</label>
        <input id="username" v-model="username" name="username">
      </p>
      <input type="submit" value="login">
    </form>
  </div>
</template>

<script>
import cookies from 'js-cookie';
import routes from '@/routes';
import axios from 'axios';

export default {
  data() {
    return { username: null };
  },
  methods: {
    login() {
      if (!this.username) {
        // TODO toast
        // eslint-disable-next-line no-alert
        window.alert('Empty username');
      }

      // vm is a dirty hack to access state from async closure ...
      const vm = this;
      axios.post(routes.postLogin, {
        username: this.username,
      }).then(() => {
        cookies.set('nienna_username', vm.username, { expires: 30 });
        vm.$store.commit('login', vm.username);
        vm.$router.push('/');
      }).catch((err) => {
        // TODO toast
        console.error('Catch error:', err);
      });
    },
  },
};
</script>
