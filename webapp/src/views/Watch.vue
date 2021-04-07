<template>
  <div class="Watch">
    <h1>This is an watch page</h1>
    <button v-on:click="play()">play</button>
    <button v-on:click="pause()">pause</button>
    <video id="video"></video>
  </div>
</template>

<script>
import Hls from 'hls.js';

export default {
  name: 'Watch',
  methods: {
    play() {
      const video = document.getElementById('video');
      video.play();
    },
    pause() {
      const video = document.getElementById('video');
      video.pause();
    },
  },
  mounted() {
    console.log('CREATED');
    if (Hls.isSupported()) {
      const video = document.getElementById('video');
      const hls = new Hls();
      hls.loadSource(`/api/videos/streams/${this.$route.params.slug}/SampleVideo_1280x720_30mb.mp4.m3u8`);
      // WANTED
      // hls.loadSource('/api/videos/streams/${VIDEOSLUG}/part.m3u8');
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, () => {
        console.log('PARSED');
        video.play();
      });
    }
  },
};
</script>
