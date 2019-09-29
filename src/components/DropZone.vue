<template>
  <div class="drop-zone" :class="{'show': show}">
    <form enctype="multipart/form-data" novalidate>
      <input type="file" @change="analyzeFile">
    </form>
    <div><i class="fas fa-cloud-upload-alt"></i></div>
    <div>
    <h1 class="selected-folder-name">Drop file to analyze</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DropZone',
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      showDropZone: false,
      filesToUpload: [],
    };
  },
  methods: {
    uploadFile(uploadObject) {
      this.$store.dispatch('media/create', uploadObject);
    },
    prepareFileForUpload(file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        this.$emit('upload', event.target.result);
      };

      reader.readAsText(file);
    },
    analyzeFile(event) {
      const files = event.target.files; //eslint-disable-line
      this.prepareFileForUpload(files[0]);
    },
  },
};
</script>

<style lang="less" scoped>
.drop-zone {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 10000;
  opacity: 0;
  visibility: hidden;
  transition: all ease 0.3s;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  input {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    height: 100%;
    width: 100%;
    opacity: 0;
  }

  &.show {
    visibility: visible;
    opacity: 1;
  }
}

i, h1 {
  font-size: 5rem;
  color: #fff;
}

h1 {
  margin-top: 2rem;
}
</style>
