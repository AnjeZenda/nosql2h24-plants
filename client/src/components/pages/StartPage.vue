<template>
  <div class="login-page">
    <div class="login-container">
      <div class="header">
        <button
            class="green-button-white-text"
            id="auth-button"
            :class="{ active: isLogin }"
            @click="switchToLogin"
        >
          Авторизация
        </button>
        <button
            class="green-button-white-text"
            id="auth-button"
            :class="{ active: !isLogin }"
            @click="switchToRegister"
        >
          Регистрация
        </button>
      </div>
      <form @submit.prevent="submitForm">
        <div v-if="!isLogin" class="form-group">
          <input
              class="inputs"
              v-model="name"
              placeholder="ФИО"
              required
          />
        </div>

        <div class="form-group">
          <input
              class="inputs"
              v-model="login"
              :placeholder="isLogin ? 'Почта/Номер телефона' : 'Почта/Номер телефона'"
              required
          />
        </div>

        <div class="form-group">
          <input
              class="inputs"
              v-model="password"
              placeholder="Пароль"
              required
          />
        </div>

        <button type="submit" class="green-button-white-text" id="login-button">
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
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      const phoneRegex = /^\+?\d{10,15}$/;
      let userData = {};

      if (emailRegex.test(this.login)) {
        userData = {
          email: this.login,
          password: this.password
        };
      } else if (phoneRegex.test(this.login)) {
        userData = {
          phone_number: this.login,
          password: this.password
        };
      }

      const response = await axios.post(AUTH_URL, userData);
      sessionStorage.setItem("id", response.data.id);
      sessionStorage.setItem("role", response.data.role);
      this.$router.push('/plants/sale');
    },

    async register() {
      const [name, surname, fatherName] = this.name.split(" ");

      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      const phoneRegex = /^\+?\d{10,15}$/;
      let userData = {};

      if (emailRegex.test(this.login)) {
        userData = {
          name: name,
          surname: surname,
          fatherName: fatherName,
          email: this.login,
          password: this.password
        };
      } else if (phoneRegex.test(this.login)) {
        userData = {
          name: name,
          surname: surname,
          fatherName: fatherName,
          phone_number: this.login,
          password: this.password
        };
      }

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
@import "../../../main.css";

.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;

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

#auth-button {
  flex: 1;
}

#auth-button.active {
  background-color: #ffffff;
  color: #000000;
}

#auth-button:not(.active):hover {
  background-color: #77934a;
}

.form-group {
  margin-bottom: 15px;
}

#login-button {
  width: 100%;
  margin-top: 10px;
}
</style>
