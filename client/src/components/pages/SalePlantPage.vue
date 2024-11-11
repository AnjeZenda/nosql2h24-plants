<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo" />
      <div class="add-form">
        <label class="create-label">Создать объявление</label>
        <form @submit.prevent="submitForm">
          <label class="label-add">
            Тип растения
            <input class="form-group-input" v-model="formData.type" type="text" placeholder="Введите тип" />
          </label>

          <label class="label-add">
            Вид растения
            <input class="form-group-input" v-model="formData.species" type="text" placeholder="Введите вид" />
          </label>

          <fieldset>
            <legend class="label-add">Размер</legend>
            <label><input v-model="formData.size" type="radio" value="Маленькие (до 20 см)" /> Маленькие (до 20 см)</label>
            <br>
            <label><input v-model="formData.size" type="radio" value="Средние (от 20 до 50 см)" /> Средние (от 20 до 50 см)</label>
            <br>
            <label><input v-model="formData.size" type="radio" value="Большие (более 50 см)" /> Большие (более 50 см)</label>
          </fieldset>

          <fieldset>
            <legend class="label-add">Условия освещения</legend>
            <label><input v-model="formData.lighting" type="radio" value="Тенелюбивые" /> Тенелюбивые</label>
            <br>
            <label><input v-model="formData.lighting" type="radio" value="Полутеневые" /> Полутеневые</label>
            <br>
            <label><input v-model="formData.lighting" type="radio" value="Светолюбивые" /> Светолюбивые</label>
          </fieldset>

          <fieldset>
            <legend class="label-add">Частота полива</legend>
            <label><input v-model="formData.wateringFrequency" type="radio" value="Редкий полив (раз в 2 недели)" /> Редкий полив (раз в 2 недели)</label>
            <br>
            <label><input v-model="formData.wateringFrequency" type="radio" value="Средний полив (1-2 раза в неделю)" /> Средний полив (1-2 раза в неделю)</label>
            <br>
            <label><input v-model="formData.wateringFrequency" type="radio" value="Частый полив (ежедневно)" /> Частый полив (ежедневно)</label>
          </fieldset>

          <fieldset>
            <legend class="label-add">Температурный режим</legend>
            <label class="radio-label"><input v-model="formData.temperature" type="radio" value="Холодостойкие (до 15°C)" /> Холодостойкие (до 15°C)</label>
            <br>
            <label><input v-model="formData.temperature" type="radio" value="Средний режим (15-22°C)" /> Средний режим (15-22°C)</label>
            <br>
            <label><input v-model="formData.temperature" type="radio" value="Теплолюбивые (более 22°C)" /> Теплолюбивые (более 22°C)</label>
          </fieldset>

          <fieldset>
            <legend class="label-add">Сложность ухода</legend>
            <label><input v-model="formData.careLevel" type="radio" value="Для начинающих" /> Для начинающих</label>
            <br>
            <label><input v-model="formData.careLevel" type="radio" value="Требует среднего ухода" /> Требует среднего ухода</label>
            <br>
            <label><input v-model="formData.careLevel" type="radio" value="Для опытных цветоводов" /> Для опытных цветоводов</label>
          </fieldset>

          <label class="label-add">
            Описание
            <textarea class="description-plant" v-model="formData.description" placeholder="Введите описание"></textarea>
          </label>

          <label class="label-add">
            Город
            <input class="form-group-input" v-model="formData.city" type="text" placeholder="Введите город" />
          </label>

          <label class="label-add">
            Цена, Р
            <input class="form-group-input" v-model="formData.price" type="number" placeholder="Введите цену" />
          </label>

          <input type="file" @change="addImage" />
          <button type="submit" class="add-button">Опубликовать</button>
        </form>
      </div>
    </div>

    <div class="plant-container">
      <div class="plant-grid">
        <div v-for="plant in plants" class="plant-card">
          <div class="plant-content">
            <img v-if="plant.image" :src="plant.image" alt="Plant Image" class="plant-image" />
            <div class="plant-info">
              <div v-if="plant.species" class="plant-title">{{ plant.species }}</div>
              <div v-if="plant.price" class="plant-price">{{ formatPrice(plant.price) }}</div>
              <div v-if="plant.place" class="plant-place">{{ plant.place }}</div>
              <div v-if="plant.createdAt" class="plant-date">{{ formatDate(plant.createdAt) }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

const PLANTS_URL = '/api/plants';
const NEW_PLANT_URL = '/api/plants/add';

export default {
  name: "Sale",
  components: { Navbar },

  data() {
    return {
      plants: [],
      formData: {
        type: '',
        species: '',
        size: '',
        lighting: '',
        wateringFrequency: '',
        temperature: '',
        careLevel: '',
        description: '',
        city: '',
        price: null,
        image: '',
      },
      userId: ''
    };
  },

  mounted() {
    this.getPlants();
    this.userId = sessionStorage.getItem("id");
  },

  methods: {
    async getPlants() {
      axios
          .get(PLANTS_URL)
          .then((response) => {
            response.data.plants.forEach(elem => {
              let plant = {
                image: elem.image,
                species: elem.species,
                price: elem.price,
                createdAt: elem.createdAt,
                place: elem.place
              };
              this.plants.push(plant)
            })
          })
    },

    addImage(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.formData.image = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },

    submitForm() {
      this.createPlant();
    },

    async createPlant() {
      const plantData = {
        image: 'https://i.pinimg.com/736x/c2/ad/d9/c2add9a552ba76ebe2c1c42e487766f7.jpg',
        place: this.formData.city,
        size: this.formData.size,
        price: this.formData.price,
        lightCondition: this.formData.lighting,
        wateringFrequency: this.formData.wateringFrequency,
        temperatureRegime: this.formData.temperature,
        careComplexity: this.formData.careLevel,
        description: this.formData.description,
        type: this.formData.type,
        species: this.formData.species,
        createdAt: new Date(),
        userId: this.userId
      };

      try {
        await axios.post(NEW_PLANT_URL, plantData);
        alert('Объявление успешно добавлено!');
        this.clearForm();
        location.reload();
      } catch (error) {
        alert('Произошла ошибка при добавлении объявления. Попробуйте снова.');
      }
    },

    clearForm() {
      this.formData.city = '';
      this.formData.size = '';
      this.formData.price = '';
      this.formData.lighting = '';
      this.formData.wateringFrequency = '';
      this.formData.temperature = '';
      this.formData.careLevel = '';
      this.formData.description = '';
      this.formData.type = '';
      this.formData.species = '';
    },

    formatDate(date) {
      return new Date(date).toLocaleDateString("ru-RU", {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      });
    },

    formatPrice(price) {
      return `${price} ₽`;
    }
  }
}
</script>

<style scoped>
.page-layout {
  display: flex;
}

.add-form {
  width: 80%;
  padding-left: 20px;
}

.logo {
  margin-top: 14%;
  margin-left: 5%;
  width: 200px;
  height: auto;
}

.create-label {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  font-weight: bold;
  display: block;
  margin-bottom: 10px;
  color: #89A758;
}

.label-add {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  font-weight: bold;
  display: block;
  margin-bottom: 10px;
}

.add-button {
  margin-top: 2%;
  width: 100%;
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  font-weight: bold;
  color: #89A758;
  background-color: transparent;
  border: 2px solid #89A758;
  border-radius: 10px;
  padding: 10px 15px;
  cursor: pointer;
  align-items: center;
  position: relative;
}

.form-group-input {
  width: 100%;
  padding-top: 10px;
  padding-bottom: 10px;
  border: 1px solid #EEECEC;
  border-radius: 10px;
  background-color: #EEECEC;
}

.plant-container {
  margin-top: 7%;
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

.plant-price {
  color: black;
  font-weight: bold;
}

.plant-place,
.plant-date {
  color: #7E7E7E;
}

textarea {
  width: 100%;
  height: 70px;
  border-radius: 8px;
  border: 1px solid #EEECEC;
  background-color: #EEECEC;
  resize: none;
}
</style>