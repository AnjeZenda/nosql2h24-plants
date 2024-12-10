<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form" style="width: 30%">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <form @submit.prevent="submitFilter">
          <div class="inputs-labels">
            Тип растения
            <input class="inputs" v-model="filter.type" type="text" placeholder="Наименование типа" />
          </div>

          <div class="inputs-labels">Условия освещения</div>
          <label class="checkbox-labels"><input v-model="filter.lightCondition" type="checkbox" value="Тенелюбивые" /> Тенелюбивые</label>
          <br>
          <label class="checkbox-labels"><input v-model="filter.lightCondition" type="checkbox" value="Полутеневые" /> Полутеневые</label>
          <br>
          <label class="checkbox-labels"><input v-model="filter.lightCondition" type="checkbox" value="Светолюбивые" /> Светолюбивые</label>

          <div class="inputs-labels">Температурный режим</div>
          <label class="checkbox-labels"><input v-model="filter.temperatureRegime" type="checkbox" value="Холодостойкие (до 15°C)" /> Холодостойкие (до 15°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="filter.temperatureRegime" type="checkbox" value="Средний режим (15-22°C)" /> Средний режим (15-22°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="filter.temperatureRegime" type="checkbox" value="Теплолюбивые (более 22°C)" /> Теплолюбивые (более 22°C)</label>

          <button type="submit" class="green-button-white-text">Отфильтровать</button>
          <button class="white-button-green-text" @click="showModal = true">Добавить информацию</button>
        </form>
      </div>
    </div>

    <div class="plant-container">
      <div class="search-plants" style="margin-bottom: 1%">
        <input class="search-input" v-model="filter.species" type="text" placeholder="Поиск растений"/>
        <button class="green-button-white-text" id="search-button" @click="submitForm">Найти</button>
      </div>

      <div class="plant-grid">
        <div v-for="(care, index) in carePlants" :key="index" class="plant-card">
          <div class="plant-content">
            <img v-if="care.image" :src="care.image" alt="Plant Image" class="plant-image" @click="getCare(care.id)"/>
            <div class="plant-info">
              <div v-if="care.species" class="plant-title">{{ care.species }}</div>
              <div v-if="care.type" class="plant-place">{{ care.type }}</div>
            </div>
          </div>
        </div>
      </div>

      <div style="display: flex; position: sticky;">
        <vue-awesome-paginate
            :total-items="careCount"
            :items-per-page="15"
            :max-pages-shown="Math.ceil(careCount / 15)"
            v-model="currentPage"
            @click="getCarePlants"
        />
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
          <input placeholder="Тип растения" type="text" class="inputs" v-model="type" />
        </div>

        <div class="form-group">
          <input placeholder="Вид растения" type="text" class="inputs"  v-model="species" />
        </div>

        <div class="form-group">
          <div class="inputs-labels">Условия освещения</div>
          <div>
            <input type="radio" value="Тенелюбивые" v-model="lightCondition" />
            <label style="font-size: 14px" >Тенелюбивые</label>
          </div>
          <div>
            <input type="radio" value="Полутеневые" v-model="lightCondition" />
            <label style="font-size: 14px" >Полутеневые</label>
          </div>
          <div>
            <input type="radio" value="Светолюбивые" v-model="lightCondition" />
            <label style="font-size: 14px" >Светолюбивые</label>
          </div>
        </div>

        <div class="form-group">
          <div class="inputs-labels">Температурный режим</div>
          <div>
            <input type="radio" value="Холодостойкие (до 15°C)" v-model="temperatureRegime" />
            <label style="font-size: 14px" >Холодостойкие (до 15°C)</label>
          </div>
          <div>
            <input type="radio" value="Средний режим (15-22°C)" v-model="temperatureRegime" />
            <label style="font-size: 14px" >Средний режим (15-22°C)</label>
          </div>
          <div>
            <input type="radio" value="Теплолюбивые (более 22°C)" v-model="temperatureRegime" />
            <label style="font-size: 14px" >Теплолюбивые (более 22°C)</label>
          </div>
        </div>

        <div class="form-group">
          <div class="inputs-labels">Введите описание ухода за растением</div>
          <textarea class="inputs" v-model="descriptionAddition"></textarea>
        </div>

        <button class="white-button-green-text" style="margin-top: -1%">
          Изображение
        </button>

        <div class="modal-footer">
          <button type="submit" class="green-button-white-text" style="margin-top: -2%">
            Опубликовать
          </button>
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
import {VueAwesomePaginate} from "vue-awesome-paginate";
import {ref} from "vue";

const CARE_PLANTS_URL = '/api/care';
const POST_CARE_URL = '/api/care/new';

export default {
  name: "Care",
  components: { VueAwesomePaginate, Navbar },

  data() {
    return {
      carePlants: [],
      showModal: false,
      filter: {
        species: '',
        type: '',
        lightCondition: [],
        temperatureRegime: []
      },
      currentPage: ref(1),
      careCount: 0,
      formData: {
        type: '',
        species: '',
        image: '',
        temperatureRegime: '',
        descriptionAddition: '',
        lightCondition: ''
      },
      type: '',
      species: '',
      image: '',
      temperatureRegime: '',
      lightCondition: '',
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

    submitFilter() {
      this.currentPage = ref(1);
      this.getPlants();
    },

    async getCarePlants() {
      this.carePlants = [];

      const careFilter = {
        filter: {
          species: this.filter.species,
          type: this.filter.type,
          lightCondition: this.filter.lightCondition,
          temperatureRegime: this.filter.temperatureRegime
        }
      }

      axios
          .post(`/api/care/${this.currentPage}/15`, careFilter)
          .then((response) => {
            response.data.plants.forEach(elem => {
              let plant = {
                image: elem.image,
                species: elem.species,
                id: elem.ruleId,
                type: elem.type
              };
              this.carePlants.push(plant)
            })
            this.careCount = parseInt(response.data.count);
          });
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
@import "../../../main.css";
@import "../../../plants.css";
@import "../../../modal.css";

#search-button {
  margin-left: 1%;
  width: 8%;
}

.author {
  color: #666;
  font-size: 0.9em;
  margin-top: 15px;
  text-align: right;
}
</style>
