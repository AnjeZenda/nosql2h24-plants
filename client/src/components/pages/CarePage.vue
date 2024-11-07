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
      care: []
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
      const response = await axios.get(CARE_PLANTS_URL);
      this.carePlants = response.data.map(elem => ({
        image: elem.image,
        species: elem.species,
        id: elem.id
      }));
    },

    async postCare() {
      const careData = {
        lightCondition: this.lightCondition,
        type: this.type,
        species: this.species,
        image: this.image,
        temperatureRegime: this.temperatureRegime,
        descriptionAddition: this.descriptionAddition,
        id: this.userId
      }
      try {
        await axios.post(POST_CARE_URL, careData);
        this.clearForm();
      } catch (error) {
        alert('Произошла ошибка при добавлении правила ухода. Попробуйте снова.');
      }
    },

    async getCare(species) {
      const response = await axios.get(`/api/care/${species}`);
      this.care = response.data.map(elem => ({
        image: elem.image,
        species: elem.species,
        id: elem.id
      }));
    },
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
</style>
