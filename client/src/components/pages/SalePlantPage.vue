<template>
  <Navbar />
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
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

const PLANTS_URL = '/api/plants';

export default {
  name: "Sale",
  components: { Navbar },

  data() {
    return {
      plants: [],
    };
  },

  mounted() {
    this.getPlants();
  },

  methods: {
    async getPlants() {
      const response = await axios.get(PLANTS_URL);
      this.plants = response.data.map(elem => ({
        image: elem.image,
        species: elem.species,
        price: elem.price,
        createdAt: elem.createdAt,
        place: elem.place
      }));
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
.plant-container {
  margin-top: 7%;
  margin-left: 20%;
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
</style>