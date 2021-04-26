<template>
  <div class="home">
    <div v-if="loading">
      <h5> loading </h5>
    </div>
    <div v-else>
      <div class="miniature_container" v-for="(video, index) in videos" :key="index">
        <Miniature :titre=video.Title :video-slug="video.Slug" />
      </div>
    </div>
  </div>
</template>

<script>
import Miniature from '@/components/Miniature.vue';
import routes from '@/routes';
import axios from 'axios';

export default {
  name: 'Home',
  components: {
    Miniature,
  },
  created() {
    this.getVideos();
  },
  data() {
    return {
      loading: false,
      videos: [],
    };
  },
  methods: {
    getVideos() {
      this.loading = true;
      axios.get(routes.getAllVideos)
        .then((response) => {
          this.loading = false;
          this.videos = response.data;
          console.log(this.videos);
        })
        .catch((err) => {
          console.log('ERROR: GETHOMEVIDEO', err);
        // TODO display error message
        });
    },
  },

};
</script>

<style>
  .miniature_container {
    display: inline;
    margin: 1%;
  }
</style>
