<!-- <template>
    <div class="container mt-5">
      <div class="d-flex justify-content-between mb-4">
        <h2>Contact List</h2>
        <div>
            <input type="text" v-model="searchQuery" class="form-control d-inline-block" style="width: 200px;" placeholder="Search by phone">
            <button class="btn btn-primary ms-2" @click="searchContacts">Search</button>
            <button class="btn btn-secondary ms-2" @click="fetchContacts">Reset</button>
        </div>
        <button class="btn btn-primary" @click="showContactForm('add')">Add Contact</button>
      </div>
  
      Contact List Table
      <table class="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Phone</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="contact in contacts" :key="contact.ID">
            <td>{{ contact.Name }}</td>
            <td>{{ contact.Phone }}</td>
            <td>
              <button class="btn btn-warning btn-sm me-2" @click="showContactForm('edit', contact)">Edit</button>
              <button class="btn btn-danger btn-sm" @click="confirmDelete(contact.ID)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
  
      Contact Form Modal
      <div class="modal fade" id="contactFormModal" tabindex="-1" aria-labelledby="contactFormModalLabel" aria-hidden="true">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="contactFormModalLabel">{{ isEditing ? 'Edit Contact' : 'Add Contact' }}</h5>
              <button type="button" class="btn-close" @click="hideContactForm" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="handleSubmit">
                <div class="mb-3">
                  <label for="contactName" class="form-label">Name</label>
                  <input type="text" v-model="contactForm.Name" class="form-control" id="contactName" required>
                </div>
                <div class="mb-3">
                  <label for="contactPhone" class="form-label">Phone</label>
                  <input type="text" v-model="contactForm.Phone" class="form-control" id="contactPhone" required>
                </div>
                <button type="submit" class="btn btn-primary">{{ isEditing ? 'Save Changes' : 'Add Contact' }}</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { Modal } from 'bootstrap';
  
  export default {
    data() {
      return {
        contacts: [], // To store the contact list
        contactForm: {
          ID: null,
          Name: '',
          Phone: ''
        },
        isEditing: false
      };
    },
    methods: {
      async fetchContacts() {
        try {
          const authToken = JSON.parse(localStorage.getItem('token'));
          const response = await fetch('http://localhost:3000/api/contact/', {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${authToken}`
            }
          });
          const data = await response.json();
  
          if (response.ok) {
            this.contacts = data; // Correctly set contacts as an array
          } else {
            
            console.error('Failed to fetch contacts:', data.error);
          }
        } catch (error) {
          console.error('Error fetching contacts:', error);
        }
      },
      async searchContacts() {
      try {
        const authToken = JSON.parse(localStorage.getItem('token'));
        const response = await fetch(`http://localhost:3000/api/contact/search?phone=${this.searchQuery}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${authToken}`
          }
        });
        const data = await response.json();

        if (response.ok) {
          this.contacts = [data]; // Set contacts as an array with the single result
        } else {
          
          console.error('Failed to fetch contact:', data.error);
        }
        } catch (error) {
            
            console.error('Error fetching contact:', error);
        }
      },
      
      showContactForm(mode, contact = {}) {
        if (mode === 'edit') {
          this.isEditing = true;
          this.contactForm = { ...contact }; // Populate form with contact data
        } else {
          this.isEditing = false;
          this.contactForm = { ID: null, Name: '', Phone: '' }; // Reset form
        }
        const contactFormModal = new Modal(document.getElementById('contactFormModal'));
        contactFormModal.show();
      },
      hideContactForm() {
        const contactFormModal = Modal.getInstance(document.getElementById('contactFormModal'));
        contactFormModal.hide();
      },
      async handleSubmit() {
        try {
          const method = this.isEditing ? 'PUT' : 'POST';
          const url = this.isEditing ? `http://localhost:3000/api/contact/${this.contactForm.ID}` : 'http://localhost:3000/api/contact/add';
          const authToken = JSON.parse(localStorage.getItem('token'));
          const response = await fetch(url, {
            method,
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${authToken}`
            },
            body: JSON.stringify(this.contactForm)
          });
          const data = await response.json();
  
          if (response.ok) {
            this.fetchContacts();
            this.hideContactForm();
          } else {
            console.error('Failed to save contact:', data.error);
          }
        } catch (error) {
          console.error('Error saving contact:', error);
        }
      },
      async confirmDelete(contactId) {
        if (confirm('Are you sure you want to delete this contact?')) {
          try {
            const authToken = JSON.parse(localStorage.getItem('token'));
            const response = await fetch(`http://localhost:3000/api/contact/${contactId}`, {
              method: 'DELETE',
              headers: {
                'Authorization': `Bearer ${authToken}`
              }
            });
            const data = await response.json();
  
            if (response.ok) {
              this.fetchContacts();
            } else {
              console.error('Failed to delete contact:', data.error);
            }
          } catch (error) {
            console.error('Error deleting contact:', error);
          }
        }
      }
    },
    mounted() {
      this.fetchContacts();
    }
  };
  </script>
  
  <style scoped>
  /* Custom styles */
  </style> -->


  <template>
    <div class="container mt-5">
        <div class="d-flex justify-content-between mb-4">
            <h2>Contact List</h2>
            <div>
                <input type="text" v-model="searchQuery" class="form-control d-inline-block" style="width: 200px;" placeholder="Search by phone">
                <button class="btn btn-primary ms-2" @click="searchContacts">Search</button>
                <button class="btn btn-secondary ms-2" @click="fetchContacts">Reset</button>
            </div>
            <button class="btn btn-primary" @click="showContactForm('add')">Add Contact</button>
        </div>

        <!-- Contact List Table -->
        <table class="table">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Phone</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="contact in contacts" :key="contact.ID">
                    <td>{{ contact.Name  }}</td>
                    <td>{{ contact.Phone  }}</td>
                    <td v-if="contact.Email">{{ contact.Email }}</td>
                    <td v-if="contact.Image"><img :src="contact.Image" alt="Profile Image" class="img-thumbnail" style="width: 50px; height: 50px;"></td>
                    <td>
                        <button class="btn btn-warning btn-sm me-2" @click="showContactForm('edit', contact)">Edit</button>
                        <button class="btn btn-danger btn-sm" @click="confirmDelete(contact.ID || contact.id)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>

        <!-- Contact Form Modal -->
        <div class="modal fade" id="contactFormModal" tabindex="-1" aria-labelledby="contactFormModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="contactFormModalLabel">{{ isEditing ? 'Edit Contact' : 'Add Contact' }}</h5>
                        <button type="button" class="btn-close" @click="hideContactForm" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <form @submit.prevent="handleSubmit">
                            <div class="mb-3">
                                <label for="contactName" class="form-label">Name</label>
                                <input type="text" v-model="contactForm.Name" class="form-control" id="contactName" required>
                            </div>
                            <div class="mb-3">
                                <label for="contactPhone" class="form-label">Phone</label>
                                <input type="text" v-model="contactForm.Phone" class="form-control" id="contactPhone" required>
                            </div>
                            <button type="submit" class="btn btn-primary">{{ isEditing ? 'Save Changes' : 'Add Contact' }}</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Modal } from 'bootstrap';

export default {
    data() {
        return {
            contacts: [], // To store the contact list
            contactForm: {
                ID: null,
                Name: '',
                Phone: ''
            },
            isEditing: false,
            searchQuery: '', // Add a searchQuery data property
        };
    },
    methods: {
        async fetchContacts() {
            try {
                const authToken = JSON.parse(localStorage.getItem('token'));
                const response = await fetch('http://localhost:3000/api/contact/', {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${authToken}`
                    }
                });
                const data = await response.json();

                if (response.ok) {
                    this.contacts = data; // Correctly set contacts as an array
                } else {
                    console.error('Failed to fetch contacts:', data.error);
                }
            } catch (error) {
                console.error('Error fetching contacts:', error);
            }
        },
        async searchContacts() {
            try {
                const authToken = JSON.parse(localStorage.getItem('token'));
                const response = await fetch(`http://localhost:3000/api/contact/search?phone=${this.searchQuery}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${authToken}`
                    }
                });
                const data = await response.json();

                if (response.ok) {
                    if (data.user) {
                        this.contacts = [data.user]; // Set contacts as an array with the single result (user profile)
                    } else if (data.name && data.phone) {
                        this.contacts = [data]; // Set contacts as an array with the single result (contact info)
                    }
                } else {
                    console.error('Failed to fetch contact:', data.error);
                }
            } catch (error) {
                console.error('Error fetching contact:', error);
            }
        },
        showContactForm(mode, contact = {}) {
            if (mode === 'edit') {
                this.isEditing = true;
                this.contactForm = { ...contact }; // Populate form with contact data
            } else {
                this.isEditing = false;
                this.contactForm = { ID: null, Name: '', Phone: '' }; // Reset form
            }
            const contactFormModal = new Modal(document.getElementById('contactFormModal'));
            contactFormModal.show();
        },
        hideContactForm() {
            const contactFormModal = Modal.getInstance(document.getElementById('contactFormModal'));
            contactFormModal.hide();
        },
        async handleSubmit() {
            try {
                const method = this.isEditing ? 'PUT' : 'POST';
                const url = this.isEditing ? `http://localhost:3000/api/contact/${this.contactForm.ID}` : 'http://localhost:3000/api/contact/add';
                const authToken = JSON.parse(localStorage.getItem('token'));
                const response = await fetch(url, {
                    method,
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${authToken}`
                    },
                    body: JSON.stringify(this.contactForm)
                });
                const data = await response.json();

                if (response.ok) {
                    this.fetchContacts();
                    this.hideContactForm();
                } else {
                    console.error('Failed to save contact:', data.error);
                }
            } catch (error) {
                console.error('Error saving contact:', error);
            }
        },
        async confirmDelete(contactId) {
            if (confirm('Are you sure you want to delete this contact?')) {
                try {
                    const authToken = JSON.parse(localStorage.getItem('token'));
                    const response = await fetch(`http://localhost:3000/api/contact/${contactId}`, {
                        method: 'DELETE',
                        headers: {
                            'Authorization': `Bearer ${authToken}`
                        }
                    });
                    const data = await response.json();

                    if (response.ok) {
                        this.fetchContacts();
                    } else {
                        console.error('Failed to delete contact:', data.error);
                    }
                } catch (error) {
                    console.error('Error deleting contact:', error);
                }
            }
        }
    },
    mounted() {
        this.fetchContacts();
    }
};
</script>

<style scoped>
/* Custom styles */
</style>

  