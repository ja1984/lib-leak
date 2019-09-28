<template>
  <div id="app">
    <div class="container">
      <div class="center">
      <div v-show="!loading">
        <textarea class="form-input form-input--large" v-model="text"></textarea>
      </div>
      <div class="button-wrapper" v-show="!loading">
        <button @click="loadData" class="button button--action">Check Dependencies</button>
      </div>
      </div>
      <div class="dependencies">
        <div class="dependency" v-for="dependency in dependencies" :key="dependency.name">
          <dependency :dependency="dependency"></dependency>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Dependency from './components/Dependency.vue';

export default {
  name: 'app',
  data() {
    return {
      text: '',
      loading: false,
      dependencies: [],
    };
  },
  components: {
    Dependency,
  },
  methods: {
    loadData() {
      this.loading = true;
      const jsonObject = JSON.parse(this.text);

      Object.keys(jsonObject.dependencies).forEach((key) => {
        this.dependencies.push({
          name: key,
          version: jsonObject.dependencies[key],
        });
      });
    },
  },
};
</script>

<style lang="less">
* {
  box-sizing: border-box;
}
#app {
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji;
  // -webkit-font-smoothing: antialiased;
  // -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}

body {
  color: #333;
  background: #f4f4f4;
}

.form-input {
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #ccc;

  &:active,
  &:focus {
    outline: none;
  }
}

.form-input--large {
  width: 100%;
  max-width: 1000px;
  height: 350px;
}

.container {
  margin: 0 auto;
  width: 90%;
  max-width: 1200;
}

.center {
  text-align: center;
}

.button {
  padding: 20px 30px;
  border-radius: 5px;
  border: none;
  background: #FF5722;
  color: #fff;
  cursor: pointer;
  transition: all ease .3s;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji;
    font-weight: bold;
    font-size: 13px;

  &:hover {
    background: lighten(#FF5722, 5%);
  }
}

.button-wrapper {
  margin: 20px;
}

.dependencies {
  margin: 0 -10px;
  display: flex;
  flex-wrap: wrap;
}

.dependency {
  padding: 10px;
  flex: 0 1 100%/3;
}
</style>
