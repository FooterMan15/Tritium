<template>
  <div>
    <nav class="navbar">
      <div class="navbar-menu">
        <router-link to="/" class="navbar-item" :class="{ active: isRouteActive('/') }">Instances</router-link>
        <router-link to="/settings" class="navbar-item" :class="{ active: isRouteActive('/settings') }">Settings</router-link>
        <div class="avatar-button" @click="navigateToAccount">
          <div class="avatar-image-container" :style="{ boxShadow: avatarBoxShadow }">
            <img src="../assets/images/wtm.png" alt="User Avatar" class="avatar-image" />
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'Navbar',
  computed: {
    avatarBoxShadow() {
      if (this.isRouteActive('/account') || this.isHoveredOverAvatar()) {
        return this.calculateBoxShadow()
      } else {
        return 'none'
      }
    }
  },
  methods: {
    calculateBoxShadow() {
      const avatarImage = document.querySelector('.avatar-image')

      const canvas = document.createElement('canvas')
      canvas.width = avatarImage.width
      canvas.height = avatarImage.height
      const ctx = canvas.getContext('2d')
      ctx.drawImage(avatarImage, 0, 0, canvas.width, canvas.height)

      const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height).data

      const [averageRed, averageGreen, averageBlue] = [...Array(3).keys()].map((channel) => {
        let sum = 0
        for (let i = channel; i < imageData.length; i += 4) {
          sum += imageData[i]
        }
        return Math.round(sum / (imageData.length / 4))
      })

      const outlineColor = `rgba(${averageRed}, ${averageGreen}, ${averageBlue}, 0.5)`

      return `0 0 10px 5px ${outlineColor}`
    },
    navigateToAccount() {
      this.$router.push('/account')
    },
    isRouteActive(route) {
      return this.$route.path === route
    },
    isHoveredOverAvatar() {
      const avatar = document.querySelector('.avatar-button')
      return avatar && avatar.matches(':hover')
    }
  }
}
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 10px;
  left: 10px;
  right: 10px;
  background-color: #f5f5f5;
  padding: 0.25rem 1rem;
  box-shadow: 0 2px 3px rgba(10, 10, 10, 0.1), 0 0 0 1px rgba(10, 10, 10, 0.1);
  z-index: 1000;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 40px;
  transition: top 0.3s ease;
}

.navbar-menu {
  display: flex;
  align-items: center;
}

.navbar-item {
  font-size: 0.9rem;
  padding: 0.25rem 0.5rem;
  text-decoration: none;
  color: #363636;
  border-radius: 20px;
  margin-right: 1rem;
  position: relative;
  transition: box-shadow 0.3s ease;
}

.navbar-item:hover {
  background-color: #e8e8e8;
}

.navbar-item.active {
  background-color: #e8e8e8;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.3);
}

.avatar-button {
  position: relative;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  cursor: pointer;
}

.avatar-image-container {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  overflow: hidden;
  display: inline-block;
}

.avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}
</style>
