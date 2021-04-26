<template>
  <div v-on:click="goto" class="miniature">
    <img :src="'/api/videos/miniature/' + this.videoSlug + '/miniature.jpeg'" alt="video miniature">
    <h6> {{ this.titre }} </h6>
  </div>
</template>
<script>

export default {
  name: 'Miniature',
  props: {
    titre: String,
    videoSlug: String,
  },
  methods: {
    goto() {
      this.$router.push({ name: 'Watch', params: { slug: this.videoSlug } });
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

  img {
    max-width: 250px;
    max-height: 150px;
  }
</style>
