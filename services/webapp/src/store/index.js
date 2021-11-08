import Vue from 'vue';
import Vuex from 'vuex';
import Cookies from 'js-cookie';

Vue.use(Vuex);

// Loading username and session states from cookies
let usernameFromCookies = '';
let loggedFromCookies = false;

if (Cookies.get('nienna')) {
  loggedFromCookies = true;
}
if (Cookies.get('nienna_username')) {
  usernameFromCookies = Cookies.get('nienna_username');
}

// Create the store
export default new Vuex.Store({
  state: {
    username: usernameFromCookies,
    logged: loggedFromCookies,
  },
  mutations: {
    toggleInit(state) {
      state.init = true;
    },
    login(state, username) {
      state.logged = true;
      state.username = username;
    },
    logout(state) {
      state.logged = false;
    },
  },
  actions: {
  },
  modules: {
  },
});
