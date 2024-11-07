<template>
  <TopAppBar />
  <div class="plant-container">
    <div class="plant-grid">
      <div v-for="(plant, index) in paginatedPlants" :key="index" class="plant-card">
        <div class="plant-content">
          <img v-if="plant.img" :src="plant.img" alt="Plant Image" class="plant-image" />
          <div class="plant-info">
            <div v-if="plant.title" class="plant-title">{{ plant.title }}</div>
            <div v-if="plant.price" class="plant-price">{{ plant.price }}</div>
            <div v-if="plant.place" class="plant-place">{{ plant.place }}</div>
            <div v-if="plant.created_at" class="plant-date">{{ plant.created_at }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="pagination">
      <button @click="prevPage" :disabled="currentPage === 1">Previous</button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
      <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
    </div>
  </div>
</template>

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

.pagination {
  margin-top: 10px;
}
</style>

<script setup>
import { ref, computed, onMounted } from 'vue';
import TopAppBar from "@/components/Navbar.vue";
import plantsData from '../../assets/plants_sell.json';

const plants = ref([]);
const currentPage = ref(1);
const pageSize = 15;

const totalPages = computed(() => {
  return Math.ceil(plants.value.length / pageSize);
});

const paginatedPlants = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  return plants.value.slice(start, start + pageSize);
});

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

onMounted(() => {
  plants.value = plantsData;
});
</script>
