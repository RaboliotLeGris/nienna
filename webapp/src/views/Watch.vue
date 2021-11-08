<template>
  <div class="Watch">
    <video id="video-player" class="video-js vjs-theme-forest" controls>
      <source :src="'/api/videos/streams/'+this.$route.params.slug+'/master.m3u8'" type="application/x-mpegURL">
    </video>
    <h4>{{this.video.title}}</h4>
    <span>{{this.video.description}}</span>
  </div>
</template>

<script>
import axios from 'axios';

import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import '@videojs/themes/dist/forest/index.css';
import qualitySelector from 'videojs-hls-quality-selector';
import qualityLevels from 'videojs-contrib-quality-levels';

import routes from '@/routes';

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
    videojs.registerPlugin('qualityLevels', qualityLevels);
    videojs.registerPlugin('hlsQualitySelector', qualitySelector);
    const player = videojs('video-player');
    player.hlsQualitySelector({
      displayCurrentQuality: true,
    });
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
