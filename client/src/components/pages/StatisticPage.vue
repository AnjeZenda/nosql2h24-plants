<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <div class="inputs-labels" id="statistic-label">
          Параметры статистики
        </div>
        <div class="inputs-labels">
          Тип статистики
        </div>
        <div class="select-sort">
          <select class="custom-select-stat" v-model="statType">
            <option disabled value="">Статистика</option>
            <option value="plants">По растениям</option>
            <option value="buy">По продажам</option>
            <option value="trade">По обменам</option>
          </select>
        </div>
        <div class="inputs-labels">
          Период времени
          <div style="display: flex; justify-content: space-between">
            <div style="display: flex; flex-direction: column; margin-right: 1%">
              <label style="font-weight: 500; font-size: 13px">С</label>
              <input class="inputs" v-model="dateFrom" type="date"/>
            </div>
            <div style="display: flex; flex-direction: column">
              <label style="font-weight: 500; font-size: 13px">По</label>
              <input class="inputs" v-model="dateTo" type="date"/>
            </div>
          </div>
        </div>
        <button style="margin-top: 2%" class="green-button-white-text" @click="getStatistic">Отобразить</button>
        <div class="inputs-labels">
          Работа с базой данных
        </div>
        <button style="margin-top: 2%" class="white-button-green-text" @click="getData">Экспортировать</button>
        <input
            type="file"
            ref="fileInput"
            @change="addDB"
            style="display: none"
        />
        <button style="margin-top: 2%" class="green-button-white-text" @click="triggerFileInput($event)">Импортировать</button>
      </div>
    </div>

    <div class="plant-container">
        <Bar v-if="chart===true" :data="data" :options="options" style="height: 500px"/>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
} from 'chart.js'
import zoomPlugin from 'chartjs-plugin-zoom';
import { Bar } from 'vue-chartjs'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend, zoomPlugin)

export default {
  name: "Statistic",
  components: { Navbar, Bar },

  data() {
    return {
      statType: '',
      dateFrom: '',
      dateTo: '',
      chart: false,
      data: {
        labels: [],
        datasets: []
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          title: {
            display: true,
            text: 'Мой график',
            font: {
              size: 18,
              weight: 'bold'
            },
            padding: {
              top: 10,
              bottom: 20
            },
            align: 'center',
            color: '#333'
          },
          legend: {
            display: true,
            position: 'top'
          },
          zoom: {
            zoom: {
              wheel: {
                enabled: true
              },
              pinch: {
                enabled: true
              },
              mode: 'x'
            },
            pan: {
              enabled: true,
              mode: 'x'
            }
          }
        },
        scales: {
          y: {
            ticks: {
              drawBorder: false,
              display: false
            },
            grid: {
              drawBorder: false,
              display: false
            }
          },
          x: {
            grid: {
              drawBorder: false
            }
          }
        }
      }
    };
  },

  methods: {
    getStatistic() {
      this.chart = true
      const datePeriod = {
        startDate: this.dateFrom,
        endDate: this.dateTo
      }
      this.data.labels = [];
      this.data.datasets = [];

      axios
          .post(`/api/statistic/${this.statType}`, datePeriod)
          .then((response) => {
            if (this.statType === 'plants') {
              const plantsCount = response.data.count;
              this.options.plugins.title.text = `Данные по опубликованным растениям. Количество записей: ${plantsCount}`;
              this.data.datasets = [
                {
                  label: 'Опубликовано',
                  backgroundColor: '#89A758',
                  data: []
                }
              ];
              response.data.statistic.forEach(elem => {
                this.data.labels.push(elem.species);
                this.data.datasets.data.push(elem.count);
              });
            } else if (this.statType === 'buy') {
              const plantsCount = response.data.count;
              this.options.plugins.title.text = `Данные по продажам. Количество записей: ${plantsCount}`;
              this.data.datasets = [
                {
                  label: 'Продано',
                  backgroundColor: '#89A758',
                  data: []
                }
              ];
              response.data.statistic.forEach(elem => {
                this.data.labels.push(elem.species);
                this.data.datasets.data.push(elem.count);
              });
            } else {
              const plantsCount = response.data.count;
              this.options.plugins.title.text = `Данные по обменам. Количество записей: ${plantsCount}`;
              let pend = [];
              let accept = [];
              let reject = [];
              response.data.statistic.forEach(elem => {
                this.data.labels.push(elem.species);
                pend.push(elem.pend);
                accept.push(elem.accept);
                reject.push(elem.reject);
              });
              this.data.datasets = [
                {
                  label: 'Создано',
                  backgroundColor: 'rgba(137,167,88,0.59)',
                  data: pend
                },
                {
                  label: 'Согласовано',
                  backgroundColor: '#a9ce6d',
                  data: accept
                },
                {
                  label: 'Отменено',
                  backgroundColor: 'rgba(137,167,88,0.11)',
                  data: reject
                },
              ];
            }
          })
    },

    async getData() {
      try {
        const response = await axios.get(`/api/data`, {});
        const jsonString = atob(response.data.db);

        const byteArray = new Uint8Array(jsonString.length);
        for (let i = 0; i < jsonString.length; i++) {
          byteArray[i] = jsonString.charCodeAt(i);
        }

        const decodedString = new TextDecoder("utf-8").decode(byteArray);

        const blob = new Blob([decodedString], { type: "application/json" });

        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.href = url;
        link.download = "data.json";
        document.body.appendChild(link);
        link.click();

        document.body.removeChild(link);
        URL.revokeObjectURL(url);
      } catch (error) {
        this.errorExport();
      }
    },

    addDB(event) {
      const reader = new FileReader();
      const file = event.target.files[0];
      reader.readAsDataURL(file);

      reader.onload = () => {
        const base64Data = reader.result.split(',')[1];

        fetch('/api/data', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            db: base64Data
          }),
        })
            .then(response => response.json())
            .then(data => {
              this.successImport();
            })
            .catch(error => {
              this.errorImport();
            });
      };

      reader.onerror = (error) => {
        console.error('Ошибка при чтении файла:', error);
      };
    },

    triggerFileInput(event) {
      event.preventDefault();
      this.$refs.fileInput.click();
    },

    errorExport() {
      this.$notify({
        title: "Ошибка!",
        text: "Произошла ошибка при импортировании базы данных.",
        type: 'error'
      });
    },

    successImport() {
      this.$notify({
        title: "Получилось!",
        text: "База данных успешно экспортирована.",
        type: 'success'
      });
    },

    errorImport() {
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

#statistic-label {
  color: #89A758;
}

.custom-select-stat {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  color: #000000;
  padding: 10px;
  border: 1px solid #EEECEC;
  background-color: #eeecec;
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  cursor: pointer;
  width: 100%;
  border-radius: 10px;
}

.custom-select-stat:focus {
  border-color: transparent;
  outline: none;
}

.custom-select-stat::after {
  position: absolute;
  border: 2px solid transparent;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
}

.custom-select-stat option {
  font-weight: 400;
  color: #000000;
  padding: 10px;
  background-color: #FFFFFF;
}
</style>