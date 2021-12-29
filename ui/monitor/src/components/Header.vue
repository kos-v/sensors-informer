<template>
  <header class="mb-3">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container-fluid">
        <a class="navbar-brand nav" href="/">{{ appName }}</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
                aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item" v-for="item in menuItems">
              <a
                  aria-current="page"
                  class="nav-link"
                  v-bind:class="{active: item.isActive}"
                  v-bind:href="item.href">{{ item.name }}</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
export default {
  name: "Header",
  data: function () {
    return {
      appName: "Sensors Monitor",
      menuItems: [
        {name: "Temperature", isActive: true, href: "#"}
      ],
    }
  }
}
</script>