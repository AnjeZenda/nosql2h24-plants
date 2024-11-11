<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo" />
      <button class="add-info-button" @click="showModal = true">Добавить информацию</button>
    </div>

    <div class="plant-container">
      <div class="plant-grid">
        <div v-for="(plant, index) in carePlants" :key="index" class="plant-card">
          <div class="plant-content">
            <img v-if="plant.image" :src="plant.image" alt="Plant Image" class="plant-image" @click="getCare(plant.species)"/>
            <div class="plant-info">
              <div v-if="plant.species" class="plant-title">{{ plant.species }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-if="showModal" class="modal-overlay" @click.self="clearForm">
    <div class="modal-content">

      <button class="close-modal" @click="clearForm">X</button>
      <h1>Информация по уходу</h1>
      <h3>Заполните основную информацию о растении</h3>

      <form @submit.prevent="submitForm">
        <div class="form-group">
          <input placeholder="Тип растения" type="text" class="form-group-input" v-model="type" />
        </div>

        <div class="form-group">
          <input placeholder="Вид растения" type="text" class="form-group-input"  v-model="species" />
        </div>

        <p>Изображение</p>
        <div class="form-group">
          <input type="file" id="imageUpload" @change="onFileChange" />
        </div>

        <div class="form-group">
          <label>Условия освещения</label>
          <div>
            <input type="radio" id="shade" value="Тенелюбивые" v-model="lightCondition" />
            <label for="shade">Тенелюбивые</label>
          </div>
          <div>
            <input type="radio" id="semiShade" value="Полутеневые" v-model="lightCondition" />
            <label for="semiShade">Полутеневые</label>
          </div>
          <div>
            <input type="radio" id="lightLoving" value="Светолюбивые" v-model="lightCondition" />
            <label for="lightLoving">Светолюбивые</label>
          </div>
        </div>

        <div class="form-group">
          <label>Температурный режим</label>
          <div>
            <input type="radio" id="coldResistant" value="Холодостойкие (до 15°C)" v-model="temperatureRegime" />
            <label for="coldResistant">Холодостойкие (до 15°C)</label>
          </div>
          <div>
            <input type="radio" id="moderate" value="Средний режим (15-22°C)" v-model="temperatureRegime" />
            <label for="moderate">Средний режим (15-22°C)</label>
          </div>
          <div>
            <input type="radio" id="heatLoving" value="Теплолюбивые (более 22°C)" v-model="temperatureRegime" />
            <label for="heatLoving">Теплолюбивые (более 22°C)</label>
          </div>
        </div>

        <div class="form-group">
          <label for="careDescription">Введите описание ухода за растением</label>
          <textarea id="careDescription" v-model="descriptionAddition"></textarea>
        </div>

        <div class="modal-footer">
          <button type="submit" class="submit-button">Опубликовать</button>
        </div>
      </form>
    </div>
  </div>

  <div v-if="isCareModalOpen" class="modal-overlay-care" @click="closeCareModal">
    <div class="modal-content-care" @click.stop>
      <header class="modal-header-care">
        <h2>{{ this.curCareSpec }}. Правила ухода.</h2>
        <button @click="closeModal" class="close-button-care">X</button>
      </header>
      <section class="modal-body-care" v-for="(care, index) in currentCare" :key="index">
        <p>{{ care.description }}</p>
        <p class="author">{{ care.author }} {{ formatDate(care.createdAt) }}</p>
      </section>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

const CARE_PLANTS_URL = '/api/care';
const POST_CARE_URL = '/api/care/new';

export default {
  name: "Care",
  components: { Navbar },

  data() {
    return {
      carePlants: [],
      showModal: false,
      lightCondition: '',
      type: '',
      species: '',
      image: '',
      temperatureRegime: '',
      descriptionAddition: '',
      userId: '',
      currentCare: [],
      curCareSpec: '',
      isCareModalOpen: false
    };
  },

  mounted() {
    this.getCarePlants();
    this.userId = sessionStorage.getItem("id");
  },

  methods: {
    submitForm() {
      this.postCare();
    },

    closeModal() {
      this.isCareModalOpen = false;
      this.currentCare = []
    },

    clearForm() {
      this.showModal = false;
      this.lightCondition = '';
      this.type = '';
      this.species = '';
      this.image = '';
      this.temperatureRegime = '';
      this.descriptionAddition = ''
    },

    async getCarePlants() {
      axios
          .get(CARE_PLANTS_URL)
          .then((response) => {
            response.data.plants.forEach(elem => {
              let plant = {
                image: elem.image,
                species: elem.species
              };
              this.carePlants.push(plant)
            })
          })
    },

    async postCare() {
      const careData = {
        species: this.species,
        descriptionAddition: this.descriptionAddition,
        userId: this.userId
      }
      try {
        await axios.post(POST_CARE_URL, careData);
        alert('Правило ухода успешно добавлено!');
        this.clearForm();
      } catch (error) {
        alert('Произошла ошибка при добавлении правила ухода. Попробуйте снова.');
      }
    },

    async getCare(species) {
      this.curCareSpec = species;
      axios
          .get(`/api/care/${species}`)
          .then((response) => {
            response.data.careRules.forEach(elem => {
              const str1 = elem.user.userSurname;
              const str2 = elem.user.userName;
              const str3 = elem.user.userFatherName;
              const result = `${str1} ${str2} ${str3}`;
              let care = {
                description: elem.description,
                createdAt: elem.createdAt,
                author: result
              };
              this.currentCare.push(care)
            })
          })
      this.isCareModalOpen = true;
    },

    formatDate(dateString) {
      const date = new Date(dateString);
      const options = {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
      };
      const formattedDate = date.toLocaleString('ru-RU', options);
      return formattedDate.replace(',', ' в');
    }
  }
}
</script>

<style scoped>
.page-layout {
  display: flex;
}

.sidebar {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 20px;
  padding-left: 10px;
}

.logo {
  margin-top: 15%;
  width: 200px;
  height: auto;
  margin-bottom: 10px;
}

.add-info-button {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  font-weight: bold;
  color: #89A758;
  background-color: transparent;
  border: 2px solid #89A758;
  border-radius: 10px;
  padding: 10px 15px;
  cursor: pointer;
  display: flex;
  align-items: center;
  position: relative;
}

.plant-container {
  margin-top: 10%;
  margin-left: 1%;
}

.plant-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.plant-card {
  width: 19%;
  box-sizing: border-box;
  margin-bottom: 30px;
}

.plant-content {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.plant-image {
  width: 170px;
  height: 170px;
  margin-bottom: 5px;
}

.plant-info {
  text-align: left;
}

.plant-title {
  color: #89A758;
  font-weight: bold;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  position: relative;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  max-width: 90%;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 15px;
}

.form-group-input {
  width: 100%;
  padding-top: 10px;
  padding-bottom: 10px;
  border: 1px solid #EEECEC;
  border-radius: 10px;
  background-color: #EEECEC;
}

.close-modal {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
  border: none;
  background: transparent;
  color: #000;
  cursor: pointer;
}

.submit-button {
  background-color: #89A758;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

textarea {
  width: 100%;
  height: 70px;
  border-radius: 8px;
  border: 1px solid #EEECEC;
  background-color: #EEECEC;
  resize: none;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

.modal-overlay-care {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content-care {
  background: #fff;
  width: 60%;
  max-width: 500px;
  padding: 20px;
  border-radius: 8px;
  position: relative;
}

.modal-header-care {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ccc;
  padding-bottom: 10px;
  margin-bottom: 10px;
}

.modal-header-care h2 {
  font-size: 1.5em;
  margin: 0;
}

.close-button-care {
  background: none;
  border: none;
  font-size: 1.5em;
  cursor: pointer;
  color: #333;
}

.modal-body-care p {
  margin: 10px 0;
  line-height: 1.6;
  font-size: 1em;
}

.author {
  color: #666;
  font-size: 0.9em;
  margin-top: 15px;
  text-align: right;
}
</style>
