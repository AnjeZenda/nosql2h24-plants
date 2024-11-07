<template>
  <div class="top-app-bar">

    <div class="menu-icons">
      <i class="icon fa fa-chart-bar"></i>
      <i class="icon fa fa-search"></i>
      <i class="icon fa fa-leaf"></i>
      <i class="icon fa fa-bell"></i>
      <i class="icon fa fa-user-circle"></i>
    </div>

    <button class="my-ads-button">Мои объявления</button>
  </div>
</template>

<script>
export default {
  name: "TopAppBar",
};
</script>

<style scoped>
.top-app-bar {
  display: flex;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
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
