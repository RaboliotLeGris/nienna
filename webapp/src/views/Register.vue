<template>
  <div class="register">
    <h1>Register</h1>
    <form @submit.prevent="register">
      <p>
        <label for="username">Username</label>
        <input id="username" v-model="username" name="username"><br>
        <label for="password">Password</label>
        <input id="password" v-model="password" name="password">
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
    return {
      username: null,
      password: null,
    };
  },
  methods: {
    register() {
      if (!this.username) {
        // TODO toast
        // eslint-disable-next-line no-alert
        window.alert('Empty username');
      }
      if (!this.password) {
        // TODO toast
        // eslint-disable-next-line no-alert
        window.alert('Empty password');
      }

      // vm is a dirty hack to access state from async closure ...
      const vm = this;
      axios.post(routes.postRegister, {
        username: this.username,
        password: this.password,
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
