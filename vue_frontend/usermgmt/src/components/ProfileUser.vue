<template>
    <div class="container mt-5">
      <!-- Profile Detail Section -->
      <div class="row">
        <div class="col-md-8 offset-md-2">
          <div class="card position-relative">
            <div class="card-body">
              <button class="btn btn-primary position-absolute top-0 end-0 mt-2 me-2" @click="showEditModal = true" data-bs-toggle="modal" data-bs-target="#editProfileModal">Edit Profile</button>
              <h3 class="card-title">Profile Details</h3>
              <div class="mb-3">
                <label for="profileName" class="form-label">Name:</label>
                <p id="profileName">{{ profile.name || 'Not provided' }}</p>
              </div>
              <div class="mb-3">
                <label for="profileEmail" class="form-label">Email:</label>
                <p id="profileEmail">{{ profile.email || 'Not provided' }}</p>
              </div>
              <div class="mb-3">
                <label for="profilePhone" class="form-label">Phone:</label>
                <p id="profilePhone">{{ profile.phone || 'Not provided' }}</p>
              </div>
              <div class="mb-3">
                <label for="profileImage" class="form-label">Profile Image:</label>
                <div v-if="profile.image">
                  <img :src="profile.image" alt="Profile Image" class="img-fluid img-thumbnail small-profile-image" />
                </div>
                <div v-else>
                  <p>No image provided</p>
                </div>
              </div>
              <div class="mb-3">
                <label for="profilePassword" class="form-label">Password:</label>
                <p id="profilePassword">{{ profile.password ? '••••••••' : 'Not provided' }}</p>
              </div>
            </div>
            <a href=""></a>
          </div>
        </div>
      </div>
  
      <!-- Edit Profile Modal -->
      <div class="modal fade" id="editProfileModal" tabindex="-1" aria-labelledby="editProfileModalLabel" aria-hidden="true" v-show="showEditModal">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editProfileModalLabel">Edit Profile</h5>
              <button type="button" class="btn-close" @click="showEditModal = false" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="handleSave">
                <div class="mb-3">
                  <label for="editName" class="form-label">Name</label>
                  <input type="text" v-model="editProfile.name" class="form-control" id="editName" required>
                </div>
                <div class="mb-3">
                  <label for="editEmail" class="form-label">Email</label>
                  <input type="email" v-model="editProfile.email" class="form-control" id="editEmail" required>
                </div>
                <div class="mb-3">
                  <label for="editPhone" class="form-label">Phone</label>
                  <input type="text" v-model="editProfile.phone" class="form-control" id="editPhone" required>
                </div>
                <div class="mb-3">
                  <label for="editImage" class="form-label">Profile Image URL</label>
                  <input type="text" v-model="editProfile.image" class="form-control" id="editImage">
                </div>
                <div class="mb-3">
                  <label for="editPassword" class="form-label">Password</label>
                  <input type="password" v-model="editProfile.password" class="form-control" id="editPassword">
                </div>
                <button type="submit" class="btn btn-primary">Save</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  
  
  export default {
    data() {
      return {
        profile: {
          name: '',
          email: '',
          phone: '',
          image: '',
          password: ''  // For display
        },
        editProfile: {
          name: '',
          email: '',
          phone: '',
          image: '',
          password: ''  // For editing
        },
        showEditModal: false
      };
    },
    methods: {  
      async fetchProfile() {
        try {
          const authToken = JSON.parse(localStorage.getItem('token'));
          const response = await fetch('http://localhost:3000/api/auth/profile', {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${authToken}`
            }
          });
          const data = await response.json();
          if (response.ok) {
            // Map response fields to component data
            this.profile = {
              name: data.Name || '',
              email: data.Email || '',
              phone: data.Phone || '',
              image: data.Image || '',
              password: data.Password || '' // Be cautious about displaying passwords
            };
            this.editProfile = { ...this.profile };
          } else {
            console.error('Failed to fetch profile:', data.error);
          }
        } catch (error) {
          console.error('Error fetching profile:', error);
        }
      },
      async handleSave() {
        try {
          const authToken = JSON.parse(localStorage.getItem('token'));
          const response = await fetch('http://localhost:3000/api/auth/profile', {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${authToken}`
            },
            body: JSON.stringify(this.editProfile)
          });
          const data = await response.json();
          if (response.ok) {
            this.profile = { ...this.editProfile };
            this.showEditModal = false;
          } else {
            console.error('Failed to update profile:', data.error);
          }
        } catch (error) {
          console.error('Error updating profile:', error);
        }
      }
    },
    mounted() {
      this.fetchProfile();
    }
  };
  </script>
  
  <style scoped>
  /* Custom styles */
  .position-absolute {
    position: absolute;
  }
  .top-0 {
    top: 0;
  }
  .end-0 {
    right: 0;
  }
  .mt-2 {
    margin-top: 0.5rem;
  }
  .me-2 {
    margin-right: 0.5rem;
  }
  .small-profile-image {
  width: 100px; /* or any desired width */
  height: 100px; /* or any desired height */
}
  </style>
  
  
  
  
  
  