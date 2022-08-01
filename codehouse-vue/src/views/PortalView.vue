<template>
  <div class="portal">
    <h1>Welcome to your Study 2gether Portal!</h1>
    <img alt="Under Construction Logo" src="../assets/portal-welcome.png" />
    <div>
        <router-link to="/"><button class="portal-button">Community Page</button></router-link>
        <router-link to="/"><button class="portal-button">Chat Page</button></router-link>
        <router-link to="/"><button class="portal-button">Niche Page</button></router-link>
        <button class="portal-button"><a href="../../public/profile-form.html">Profile Page</a></button>
    </div>
    <router-link to="/"><button>Logout</button></router-link>
  </div>
  
</template>

<style scoped>

.portal-button{
  background-color: #42273B;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
}
button {
  background-color: #3A405A;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
}
a {
    text-decoration: none;
    color: white;
}
</style>