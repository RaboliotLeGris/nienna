<template>
  <div class="container">
    <div v-show="this.$store.state.logged" class="large-12 medium-12 small-12 cell">
      <div v-if="this.status === 'None'">
        <label> Title
          <input type="text" id="title" ref="title" v-on:change="handleFileTitle()">
        </label><br>
        <label>File
          <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
        </label>
        <button v-on:click="submitFile()">Submit</button>
      </div>
      <div v-else>
        <h5>{{ this.status }}</h5>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import routes from '@/routes';

export default {
  data() {
    return {
      status: 'None',
      title: '',
      file: '',
    };
  },
  methods: {
    submitFile() {
      // TODO: improve this bits
      this.status = 'Uploading';
      const formData = new FormData();
      console.log('TITLE', this.title);
      formData.append('title', this.title);
      formData.append('video', this.file);
      axios.post(routes.postVideo,
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        })
        .then((res) => {
          this.status = 'Uploaded';
          setInterval(() => {
            this.pollStatus(res.data.Slug);
          }, 5000);
        })
        .catch((err) => {
          console.log('FAILURE!!', err);
        });
    },
    handleFileUpload() {
      // eslint-disable-next-line prefer-destructuring
      this.file = this.$refs.file.files[0];
    },
    handleFileTitle() {
      this.title = this.$refs.title.value;
    },
    pollStatus(slug) {
      axios.get(routes.getStatusVideo + slug)
        .then((response) => {
          this.status = response.data.status;
        })
        .catch((err) => {
          console.log('ERROR: GETSTATUSVIDEO', err);
          // TODO display error message
        });
    },
  },
};
</script>
