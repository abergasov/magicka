<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <MainPage msg="Welcome to magick"/>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import MainPage from './components/Main.vue';

export default defineComponent({
  name: 'App',
  components: {
    MainPage
  }
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
