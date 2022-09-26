<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button @click="connectWallet">Connect wallet</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type {AxiosInstance} from 'axios'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance
  }
}

export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  methods: {
    async connectWallet() {
      const tagsResponse = await this.$axios.get('/ping');
      console.log('connect wallet', tagsResponse);
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
