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
              v-model="fullName"
              placeholder="ФИО"
              required
          />
        </div>

        <div class="form-group">
          <input
              type="text"
              v-model="username"
              :placeholder="isLogin ? 'Почта/Номер телефона' : 'Почта/номер телефона'"
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
export default {
  data() {
    return {
      isLogin: true,
      username: '',
      password: '',
      fullName: '',
    };
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
      this.username = '';
      this.password = '';
      this.fullName = '';
    },
    submitForm() {
      if (this.isLogin) {
        console.log('Авторизация:', this.username, this.password);
      } else {
        console.log('Регистрация:', this.fullName, this.username, this.password);
      }
    },
  },
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
  background-color: #fff;
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
  font-size: 18px;
  font-weight: bold;
  background-color: #89A758;
  color: #fff;
  padding: 10px;
  border: none;
  border-radius: 10px 10px 0 0;
  cursor: pointer;
}

.auth-button.active {
  background-color: #ffffff;
  color: #000000;
  border-bottom: 2px solid #89A758;
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
  border: 1px solid #ddd;
  border-radius: 10px;
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
}

.login-button:hover {
  background-color: #77934a;
}
</style>
