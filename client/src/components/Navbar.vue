<template>
  <div class="navbar">

    <div class="menu-icons">
      <i v-if="role === 0" class="icon fa fa-chart-bar" @click="navigate(0)"></i>
      <i class="icon fa fa-search" @click="navigate(role === 0 ? 1 : 0)"></i>
      <i class="icon fa fa-leaf" @click="navigate(role === 0 ? 2 : 1)"></i>
      <i class="icon fa fa-bell"></i>
      <i class="icon fa fa-user-circle" @click="navigate(role === 0 ? 3 : 2)"></i>
    </div>

    <button class="my-ads-button" @click="navigate(role === 0 ? 4 : 3)">Мои объявления</button>
  </div>
</template>

<script>
export default {
  name: "Navbar",

  data() {
    return {
      role: '',
      adminItems: [
        { path: '/plants/statistics' },
        { path: '/plants/sale' },
        { path: '/plants/care'},
        { path: '/plants/user' },
        { path: '/plants/ads' }
      ],
      regularItems: [
        { path: '/plants/sale' },
        { path: '/plants/care'},
        { path: '/plants/user' },
        { path: '/plants/ads' }
      ]
    };
  },

  beforeMount() {
    this.role = sessionStorage.getItem("role");
  },

  computed: {
    menuItems() {
      return this.role === 0 ? this.adminItems : this.regularItems;
    }
  },

  methods: {
    navigate(index) {
      const target = this.menuItems[index];
      if (target) {
        this.$router.push(target.path);
      }
    }
  }
};
</script>

<style scoped>
.navbar {
  display: flex;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  position: fixed;
  top: 0;
  left: 0;
  width: 98%;
  z-index: 1000;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

.menu-icons {
  display: flex;
  gap: 15px;
  align-items: center;
  margin-left: auto;
}

.icon {
  font-size: 18px;
  cursor: pointer;
  color: #fff;
}

.my-ads-button {
  padding: 8px 16px;
  background-color: #89A758;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  margin-left: 20px;
}

.my-ads-button:hover {
  background-color: #77934a;
}
</style>
