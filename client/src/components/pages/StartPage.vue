<template>
  <div class="login-page">
    <div class="login-container">
      <div class="header">
        <button
            class="auth-button"
            :class="{ active: isLogin }"
            @click="switchToLogin"
        >
          Авторизация
        </button>
        <button
            class="auth-button"
            :class="{ active: !isLogin }"
            @click="switchToRegister"
        >
          Регистрация
        </button>
      </div>
      <form @submit.prevent="submitForm">
        <div v-if="!isLogin" class="form-group">
          <input
              type="text"
              v-model="name"
              placeholder="ФИО"
              required
          />
        </div>

        <div class="form-group">
          <input
              type="text"
              v-model="login"
              :placeholder="isLogin ? 'Почта/Номер телефона' : 'Почта/Номер телефона'"
              required
          />
        </div>

        <div class="form-group">
          <input
              type="password"
              v-model="password"
              placeholder="Пароль"
              required
          />
        </div>

        <button type="submit" class="login-button">
          {{ isLogin ? 'Войти' : 'Зарегистрироваться' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

const REGISTER_URL = '/api/v1/register'
const AUTH_URL = '/api/v1/login'

export default {
  data() {
    return {
      isLogin: true,
      login: '',
      password: '',
      name: '',
    };
  },

  beforeMount() {
    window.sessionStorage.clear();
  },

  methods: {
    switchToLogin() {
      this.isLogin = true;
      this.clearForm();
    },

    switchToRegister() {
      this.isLogin = false;
      this.clearForm();
    },

    clearForm() {
      this.login = '';
      this.password = '';
      this.name = '';
    },

    submitForm() {
      if (this.isLogin) {
        this.auth();
      } else {
        this.register();
      }
    },

    async auth() {
      const userData = {
          login: this.login,
          password: this.password
      };

      const response = await axios.post(AUTH_URL, userData);
      sessionStorage.setItem("id", response.data.id);
      sessionStorage.setItem("role", response.data.role);
      this.$router.push('/plants/sale');
    },

    async register() {
      const [name, surname, fatherName] = this.name.split(" ");
      const userData = {
        name: name,
        surname: surname,
        fatherName: fatherName,
        email: this.login,
        password: this.password
      };

      try {
        await axios.post(REGISTER_URL, userData);
        alert('Регистрация выполнена успешно!');
      } catch (error) {
        alert('Произошла ошибка при регистрации. Попробуйте снова.');
      }
    }
  }
};
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  font-family: 'Century Gothic', sans-serif;
}

.login-container {
  width: 400px;
  padding: 20px;
  background-color: #FFFFFF;
  border: 1px solid #989898;
  border-radius: 10px;
  box-sizing: border-box;
}

.header {
  display: flex;
  margin-bottom: 20px;
}

.auth-button {
  flex: 1;
  font-size: 16px;
  font-weight: bold;
  background-color: #89A758;
  color: #fff;
  padding: 10px;
  border: none;
  border-radius: 10px;
  cursor: pointer;
}

.auth-button.active {
  background-color: #ffffff;
  color: #000000;
}

.auth-button:not(.active):hover {
  background-color: #77934a;
}

.form-group {
  margin-bottom: 15px;
}

input[type="text"],
input[type="password"] {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  border-radius: 10px;
  border: none;
  color: #676767;
  background-color: #EDEDED;
  box-sizing: border-box;
}

.login-button {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  background-color: #89A758;
  color: #fff;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  margin-top: 10px;
  font-weight: bold;
}

.login-button:hover {
  background-color: #77934a;
}
</style>
