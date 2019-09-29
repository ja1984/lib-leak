<template>
  <div id="app" @dragenter="toggleDropZone" @drop="toggleDropZone" @dragleave="toggleDropZone">
    <section class="search-wrapper" :class="{'search-wrapper--center': dependencies.length === 0}">
      <div class="container">
        <h1 class="logo">lib/leak</h1>
        <form @submit.prevent="loadData">
          <div>
            <div class="input-wrapper">
              <input class="form-input" placeholder="NMP Package" type="text" v-model="text" />
            </div>
            <p
              class="pro-tip"
            >PRO Tip, you can drag en drop your package.json to check all dependencies</p>
            <!-- <textarea class="form-input form-input--large" v-model="text"></textarea> -->
            <div class="center">
              <div class="button-wrapper">
                <button type="submit" class="button button--action">Check Dependencies</button>
              </div>
            </div>
          </div>
        </form>
      </div>
            <div class="container">
        <div class="dependencies">
          <div class="dependency" v-for="dependency in dependencies" :key="dependency.name">
            <dependency :dependency="dependency"></dependency>
          </div>
        </div>
      </div>
          <section class="toolbar">
      <span>
        Made with
        <heart-icon size="1.5x" class="toolbar-icon toolbar-icon--red"></heart-icon>in Varberg
      </span>
      /
      <span>
        <a class="toolbar__link" href="https://github.com/ja1984/lib-leak" target="blank">
          <github-icon size="1.5x" class="toolbar-icon"></github-icon>Github
        </a>
      </span>
      /
      <span>
        <a class="toolbar__link" href="https://github.com/ja1984/lib-leak" target="blank">
          <github-icon size="1.5x" class="toolbar-icon"></github-icon>Github
        </a>
      </span>
    </section>
    </section>


    <drop-zone :show="showDropZone" @upload="upload"></drop-zone>
  </div>
</template>

<script>
import { HeartIcon, GithubIcon } from 'vue-feather-icons';
import DropZone from '@/components/DropZone.vue';
import Dependency from './components/Dependency.vue';

export default {
  name: 'app',
  data() {
    return {
      text: '',
      json: '',
      loading: false,
      dependencies: [],
      showDropZone: false,
    };
  },
  components: {
    Dependency,
    HeartIcon,
    DropZone,
    GithubIcon,
  },
  methods: {
    toggleDropZone() {
      this.showDropZone = !this.showDropZone;
    },
    upload(json) {
      const jsonObject = JSON.parse(json);
      Object.keys(jsonObject.dependencies).forEach((key) => {
        this.dependencies.push({
          name: key,
          version: jsonObject.dependencies[key],
        });
      });
    },
    loadData() {
      this.loading = true;
      this.dependencies.push({
        name: this.text,
        version: '',
      });
      this.text = '';
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
  padding: 15px;
  border: none;
  width: 100%;
  border-radius: 6px;
  font-size: 20px;

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
  background: #ff5722;
  color: #fff;
  cursor: pointer;
  transition: all ease 0.3s;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji;
  font-weight: bold;
  font-size: 13px;
  white-space: nowrap;

  &:hover {
    background: lighten(#ff5722, 5%);
  }
}

.button-wrapper {
  margin: 20px;
}

.dependencies {
  margin: 0 -10px;
  display: flex;
  flex-wrap: wrap;
  padding-bottom: 50px;
}

.dependency {
  padding: 10px;
  flex: 0 1 100%/3;
}

.logo {
  font-family: "Raleway", sans-serif;
  font-size: 60px;
  text-align: center;
  margin: 30px 0;
}

.input-wrapper {
  position: relative;
  border-radius: 6px;
  background: #fff;
  box-shadow: 0 5px 10px 0 rgba(168, 182, 191, 0.6);
  width: 100%;
  max-width: 700px;
  margin: 0 auto;
}

.search-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  display: flex;
  align-items: center;
  flex-direction: column;
  overflow: auto;
  transition: all ease .3s;
}

.search-wrapper--center {
  justify-content: center;
}

.toolbar {
  position: absolute;
  top: 20px;
  right: 20px;
  color: #aaa;
  font-size: 13px;
}

.toolbar__link {
  color: #aaa;
  text-decoration: none;
}
.toolbar-icon {
  vertical-align: middle;
}
.toolbar-icon--red {
  color: #f44336;
}

.pro-tip {
  font-size: 13px;
  text-align: center;
  color: #ccc;
}
</style>
