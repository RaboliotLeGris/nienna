<template>
  <div class="Watch">
    <video id="video-player" class="video-js vjs-theme-forest" controls>
      <source :src="'/api/videos/streams/'+this.$route.params.slug+'/part.m3u8'" type="application/x-mpegURL">
    </video>
    <h4>{{this.video.Title}}</h4>
    <span>{{this.video.Description}}</span>
  </div>
</template>

<script>
import routes from '@/routes';
import axios from 'axios';

import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import '@videojs/themes/dist/forest/index.css';

export default {
  name: 'Watch',
  data() {
    return {
      video: {},
    };
  },
  created() {
    this.getVideos();
  },
  mounted() {
    videojs('video-player');
  },
  methods: {
    getVideos() {
      axios.get(routes.getOneVideo + this.$route.params.slug)
        .then((response) => {
          this.video = response.data;
        })
        .catch((err) => {
          console.log('ERROR: GETONEVIDEO', err);
          // TODO display error message
        });
    },
  },
};
</script>

<style>
#video-player {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
