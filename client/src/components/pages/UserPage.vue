<template>
  <Navbar/>
  <div class="sidebar">
    <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo" />
    <div class="user-info">
      <img :src="user.avatar" alt="User Photo" class="user-photo">
      <div class="user-details">
        <p class="user-lastname">{{ user.lastName }}</p>
        <p class="user-firstname">{{ user.firstName }}</p>
        <p class="user-middlename">{{ user.patronymic }}</p>
      </div>
    </div>
    <p class="create">Cоздание аккаунта: {{ formatDate(user.accountCreationDate) }}</p>
    <p class="create">Внесение изменений: {{ formatDate(user.lastModificationDate) }}</p>
    <div class="add-form">
      <form>
        <label class="label-add"> Фамилия
          <input class="form-group-input" v-model="user.lastName" type="text" :placeholder="user.lastName" />
        </label>

        <label class="label-add"> Имя
          <input class="form-group-input" v-model="user.firstName" type="text" :placeholder="user.firstName" />
        </label>

        <label class="label-add"> Отчество
          <input class="form-group-input" v-model="user.patronymic" type="text" :placeholder="user.patronymic" />
        </label>

        <label class="label-add"> Почта
          <input class="form-group-input" v-model="user.email" type="text" :placeholder="user.email" />
        </label>

        <label class="label-add"> Телефон
          <input class="form-group-input" v-model="user.phone" type="text" :placeholder="user.phone" />
        </label>

        <button class="save-button">Сохранить изменения</button>
        <button class="out-button">Выход</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import Navbar from "@/components/Navbar.vue";
export default {
  name: "User",
  components: { Navbar },

  data() {
    return {
      user: {
        avatar: '',
        accountCreationDate: '',
        lastModificationDate: '',
        lastName: '',
        firstName: '',
        patronymic: '',
        email: '',
        phone: '',
        userId: ''
      }
    };
  },

  mounted() {
    this.user.userId = sessionStorage.getItem("id");
    this.getUser();
  },

  methods: {
    async getUser() {
      const id = this.user.userId;
      console.log(id);
      axios
          .get(`/api/v1/user/${id}`)
          .then((response) => {
            const curUser = response.data;
            console.log(curUser);
            this.user.avatar = curUser.photo;
            this.user.lastName = curUser.surname;
            this.user.firstName = curUser.name;
            this.user.patronymic = curUser.fatherName;
            this.user.email = curUser.email;
            this.user.phone = curUser.phoneNumber;
            this.user.lastModificationDate = curUser.updatedAt;
            this.user.accountCreationDate = curUser.createdAt;
          })
    },

    formatDate(dateString) {
      const date = new Date(dateString);
      const options = {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
      };
      return date.toLocaleDateString('ru-RU', options);
    }
  }
};
</script>


<style scoped>
.user-info {
  padding-left: 20px;
  display: flex;
  width: 18%;
  justify-content: space-between;
  align-items: flex-start; /* Выравнивание по верхнему краю */
}

.user-photo {
  width: 80px;  /* Задайте нужный размер фото */
  height: 80px;
  border-radius: 50%; /* Сделать фото круглым */
  margin-right: 10px; /* Расстояние между фото и текстом */
}

.user-details {
  display: flex;
  flex-direction: column; /* Размещение текста в столбик */
}

.user-details p {
  text-align: right;
  margin: 0;
  font-family: 'Century Gothic', sans-serif;
  font-size: 14px;
  font-weight: bold;
  /* Убираем отступы между строками */
}

.add-form {
  width: 18%;
  padding-left: 20px;
}

.logo {
  margin-top: 3%;
  margin-left: 5%;
  width: 200px;
  height: auto;
}

.label-add {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  font-weight: bold;
  display: block;
  margin-bottom: 10px;
}

.form-group-input {
  width: 100%;
  padding-top: 10px;
  padding-bottom: 10px;
  border: 1px solid #EEECEC;
  border-radius: 10px;
  background-color: #EEECEC;
}

.save-button {
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

.out-button {
  margin-top: 2%;
  width: 100%;
  font-family: 'Century Gothic', sans-serif;
  background-color: #89A758;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-weight: bold;
}

.create {
  color: #474747;
  font-size: 1em;
  padding-left: 20px;
  margin-top: 1px;
}
</style>