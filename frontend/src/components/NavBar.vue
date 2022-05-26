<template>
  <div class="nav-bar">
    <div class="logo">
      <a href="">
        <img
          src="https://aura.network/assets/imgs/auraMainLogo-1x.png"
          alt=""
        />
      </a>
    </div>
    <div class="nav">
      <ul>
        <li><a href="">All events</a></li>
        <li><a href="">By Validator</a></li>
      </ul>
    </div>
    <div class="link">
      <ul>
        <li>
          <a href=""><i class="fa-brands fa-instagram"></i></a>
        </li>
        <li>
          <a href=""><i class="fa-brands fa-facebook"></i></a>
        </li>
        <li>
          <a href=""><i class="fa-brands fa-linkedin"></i></a>
        </li>
        <li>
          <a href=""><i class="fa-brands fa-telegram"></i></a>
        </li>
      </ul>
    </div>
    <div class="login">
      <div class="content">Login</div>
    </div>
  </div>
</template>

<script>
export default {
  setup() {},
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Anton&display=swap");
.nav-bar {
  display: flex;
  background: #141416;
  font-family: "Anton", sans-serif !important;
  color: white;
  justify-content: center;
  align-items: center;
  padding: 10px;
  font-size: 20px;
}
.nav ul {
  list-style: none;
  /* display: flex; */
}
.nav ul li {
  display: inline;
  margin-right: 40px;
}
.link ul li {
  display: inline;
  margin-right: 20px;
}
ul li a {
  text-decoration: none;
  color: white;
}
.logo {
  flex-grow: 1;
}
.nav {
  flex-grow: 12;
}
.login {
  flex-grow: 1;
  text-align: center;
  margin-right: 10px;
  padding: 2px 2px;
  background: linear-gradient(to right, red, yellow);
  border-radius: 10px;
}
.login .content {
  width: 100%;
  height: 100%;
  background: #141416;
  border-radius: 10px;
  cursor: pointer;
}
.nav a:hover {
  color: mistyrose;
}
.link a:hover {
  font-size: 30px;
  transition: 0.2s ease-in-out;
}
</style>
