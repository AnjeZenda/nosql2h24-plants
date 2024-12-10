<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <div class="inputs-labels" id="create-ads">
          Работа с базой данных
        </div>
        <button style="margin-top: 2%" class="white-button-green-text" @click="getData">Импортировать</button>
        <input
            type="file"
            ref="fileInput"
            @change="addDB"
            style="display: none"
        />
        <button style="margin-top: 2%" class="green-button-white-text">Экспортировать</button>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

export default {
  name: "Statistic",
  components: { Navbar },

  data() {
    return {
      data: ''
    };
  },

  methods: {
    async getData() {
      try {
        const response = await axios.get(`/api/data`, {
          responseType: 'blob'
        });
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;

        link.setAttribute('importDB', 'data.json');
        document.body.appendChild(link);
        link.click();

        document.body.removeChild(link);
        this.successImport();
      } catch (error) {
        this.errorImport();
      }
    },

    addDB(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.data = e.target.result;
        };
        reader.readAsDataURL(file);
        this.postData();
      }
    },

    triggerFileInput(event) {
      event.preventDefault();
      this.$refs.fileInput.click();
    },

    async postData() {
      try {
        await axios.post(`/api/data`, this.data);
        this.successExport();
      } catch (error) {
        this.errorExport();
      }
    },

    successImport() {
      this.$notify({
        title: "Получилось!",
        text: "База данных успешно импортирована.",
        type: 'success'
      });
    },

    errorImport() {
      this.$notify({
        title: "Ошибка!",
        text: "Произошла ошибка при импортировании базы данных.",
        type: 'error'
      });
    },

    successExport() {
      this.$notify({
        title: "Получилось!",
        text: "База данных успешно экспортирована.",
        type: 'success'
      });
    },

    errorExport() {
      this.$notify({
        title: "Ошибка!",
        text: "Произошла ошибка при экспортировании базы данных.",
        type: 'error'
      });
    },
  }
}
</script>

<style scoped>
@import "../../../main.css";
@import "../../../plants.css";

.page-layout {
  display: flex;
}

#create-ads {
  color: #000000;
}
</style>