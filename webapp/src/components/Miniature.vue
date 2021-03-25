<template>
  <div v-on:click="goto" class="miniature">
<!--    <img :src="miniatureURL" alt="video miniature">-->
    <img src="@/assets/miniature-placeholder.png" alt="video miniature">
    <h6> {{ this.titre }} </h6>
  </div>
</template>
<!-- @/assets/miniature-placeholder.png -->
<script>

export default {
  name: 'Miniature',
  props: {
    titre: String,
    miniatureURL: String,
    videoSlug: String,
  },
  methods: {
    goto() {
      this.$router.push(`/view/${this.videoSlug}`);
    },
  },
};

</script>

<style>
  .miniature {
    padding: 5px;
    background-color: #42b983;
    width: 15%;
    display: inline-block;
  }
</style>
