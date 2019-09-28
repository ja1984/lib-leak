<template>
  <div class="dependency">
    <header class="dependency__header">
          <div class="dependency__item dependency__item--fill">
      <div class="version">{{ dependency.version }}</div>
      <strong>{{ dependency.name }}</strong>
    </div>
    <div class="dependency__item">
      <div class="status">
        <loader-icon size="1.5x" class="custom-class icon loading" :class="{'icon--hide': !loading}"></loader-icon>
        <check-icon size="1.5x" class="custom-class icon icon--success" :class="{'icon--hide': error || loading || references.length > 0}"></check-icon>
        <alert-circle-icon size="1.5x" class="custom-class icon icon--warning" :class="{'icon--hide': loading || !error}"></alert-circle-icon>
      <alert-triangle-icon size="1.5x" class="custom-class icon icon--error" v-if="references.length > 0"></alert-triangle-icon>
      </div>
    </div>
    </header>
    <section class="dependency__body">
      <ul class="references" v-if="references.length > 0">
        <li v-for="issue in references" :key="issue.id" class="reference__item">
            <lock-icon size="1x" v-if="issue.state !== 'open'" class="custom-class state-icon"></lock-icon>
            <unlock-icon size="1x" v-if="issue.state === 'open'" class="custom-class state-icon"></unlock-icon>
            <a :href="issue.html_url" target="_blank">{{ issue.title }}</a>
        </li>
      </ul>
    </section>

  </div>
</template>

<script>
import axios from 'axios';
import {
  LoaderIcon, AlertTriangleIcon, LockIcon, UnlockIcon, CheckIcon, AlertCircleIcon,
} from 'vue-feather-icons';

export default {
  name: 'Dependency',
  data() {
    return {
      loading: true,
      error: false,
      leak: false,
      references: [],
    };
  },
  components: {
    LoaderIcon,
    AlertTriangleIcon,
    LockIcon,
    UnlockIcon,
    CheckIcon,
    AlertCircleIcon,
  },
  mounted() {
    axios.get(`https://api.npms.io/v2/package/${this.dependency.name}`).then((response) => {
      const repositoryUrl = response.data.collected.metadata.links.repository || response.data.collected.metadata.links.homepage;
      if (!repositoryUrl) {
        this.error = true;
        this.loading = false;
        return;
      }
      this.checkIssues(repositoryUrl);
      // this.checkReleases(repositoryUrl);
    }).catch((error) => {
      console.log(error);
      this.error = true;
      this.loading = false;
    });
  },
  methods: {
    containsMemoryLeak(string) {
      const _string = string.toLowerCase();
      if (_string.includes('memoryleak')) return true;
      if (_string.includes('memory leak')) return true;
      return false;
    },
    checkIssues(url) {
      const issuesUrl = `https://api.github.com/repos${url.split('https://github.com')[1]}/issues`;
      axios.get(issuesUrl).then((response) => {
        this.loading = false;

        response.data.forEach((issue) => {
          if (this.containsMemoryLeak(issue.body) || this.containsMemoryLeak(issue.title)) {
            this.references.push(issue);
          }
        });
      });
    },
    checkReleases(url) {
      const releasesUrl = `https://api.github.com/repos${url.split('https://github.com')[1]}/releases`;
      axios.get(releasesUrl).then((response) => {
        console.log(response.data);
      });
    },
  },
  props: {
    dependency: {
      type: Object,
      default: () => { },
    },
  },
};
</script>

<style lang="less" scoped>
.dependency {
  border-radius: 6px;
  box-shadow: 0 5px 10px 0 rgba(168, 182, 191, 0.6);
  transition: all ease 0.3s;
  overflow: hidden;
  background: #fff;
  // margin-bottom: 30px;
  padding: 15px;
}

ul {
  margin: 0;
  padding: 0;
  list-style-type: none;
}

.dependency__header {
  display: flex;
  align-items: center;
  margin: 0 -10px;

}
.dependency__item {
  padding: 0 10px;
}
.dependency__item--fill {
  flex: 1;
  padding: 0 10px;
}

.version {
  color: #ccc;
  font-size: 12px;
}

.custom-class {
  display: inline-block;
  vertical-align: middle;
}

a {
  color: #666;
  transition: all ease .3s;

  &:hover {
    color: #333;
  }
}

@-webkit-keyframes rotating {
  from {
    -webkit-transform: rotate(0deg);
  }
  to {
    -webkit-transform: rotate(360deg);
  }
}

.rotating {
  -webkit-animation: rotating 2s linear infinite;
}

.icon {
  transition: all ease 0.3s;
}

.icon--hide {
  opacity: 0;
  visibility: hidden;
}

.icon--error {
  color: #F44336;
}

.icon--success {
  color: #8BC34A;
}
.icon--warning {
  color: #FFC107;
}

.status {
  position: relative;
  height: 24px;
  width: 24px;

  .icon {
    position: absolute;
    top: 0;
    left: 0;
  }
}

.state-icon {
  margin-right: 10px;
}

.reference__item {
  padding: 5px 0;
}

.references {
  margin-top: 10px;
}
</style>
