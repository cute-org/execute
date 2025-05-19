<template>
  <Dialog as="div" :open="isOpen" @close="closeDialog" class="relative z-10">
    <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <DialogPanel class="bg-infoBg rounded-lg p-6 w-full max-w-md relative border-borderColor border-2">
        <!-- Close Button -->
        <button 
          @click="closeDialog"
          class="absolute top-4 right-4 text-gray-500 hover:text-gray-300"
        >
          <span class="text-xl font-bold">X</span>
        </button>

        <!-- Dialog Content -->
        <div class="space-y-6">
          <h2 class="text-2xl font-bold text-center text-white">Edit Settings</h2>
            <!-- Avatar Upload Section -->
          <div class="flex flex-col items-center">
            <div class="w-24 h-24 mb-3 relative">
              <img 
                :src="avatarPreview || defaultAvatar" 
                alt="Avatar" 
                class="w-full h-full rounded-full object-cover border-2 border-gray-600"
              />
              <button 
                @click="triggerAvatarInput"
                class="absolute bottom-0 right-0 bg-blue-600 rounded-full p-1 text-white"
                title="Change avatar"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </button>
            </div>
            <input 
              type="file" 
              ref="avatarInput" 
              @change="handleFileChange" 
              accept="image/*" 
              class="hidden" 
            />
            <button 
              v-if="avatarPreview" 
              @click="removeAvatar" 
              class="text-sm text-gray-400 hover:text-gray-300"
            >
              Remove avatar
            </button>
          </div>
            <!-- Username Field -->
            <div>
              <label for="displayname" class="block text-sm font-medium text-gray-300 mb-1">
                Display name
              </label>
              <input
                id="displayname"
                type="text"
                v-model="displayname"
                class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Enter your username"
                required
              />
            </div>
            <!-- Role Field -->
            <div>
              <label for="role" class="block text-sm font-medium text-gray-300 mb-1">
                Role
              </label>
              <input
                id="role"
                type="text"
                v-model="role"
                class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Enter your role"
                required
              />
            </div>

            <!-- Phone Field -->
            <div>
              <label for="phone" class="block text-sm font-medium text-gray-300 mb-1">
                Phone Number
              </label>
              <input
                id="phone"
                type="tel"
                v-model="phone"
                class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="(+xx)xxx-xxx-xxx"
              />
            </div>
            
            <!-- Birthdate Field -->
            <div>
              <label for="birthdate" class="block text-sm font-medium text-gray-300 mb-1">
                Birthdate
              </label>
              <input
                id="birthdate"
                type="date"
                v-model="birthdate"
                class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <!-- Password verification field -->
            <div>
              <label for="password" class="block text-sm font-medium text-gray-300 mb-1">
                Password
              </label>
              <input
                id="password"
                type="password"
                v-model="password"
                placeholder = "Enter your password to submit changes"
                class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>

            <div class="flex justify-center items-center">
              <!-- Error Message -->
              <p v-if="error" class="text-error text-sm">{{ error }}</p>
              <!-- Success Message -->
              <p v-if="success" class="text-accepted text-sm">{{ success }}</p>
            </div>
            


            <!-- Buttons -->
            <div class="flex gap-3">
              <button
                type="button"
                @click="closeDialog"
                class="flex-1 py-2 px-4 border border-gray-600 rounded-md text-gray-300 hover:bg-gray-700 transition-colors"
              >
                Cancel
              </button>

              <button
                type="submit"
                @click="handleSubmit"
                :disabled="isLoading"
                class="flex-1 py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50"
              >
                {{ isLoading ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
        </div>
      </DialogPanel>
    </div>
  </Dialog>
</template>

<script setup>
import { ref, watch } from 'vue'
import { Dialog, DialogPanel } from '@headlessui/vue'
import { addISOWeekYears } from 'date-fns'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:show'])

const isOpen = ref(props.show)
const displayname = ref('')
const birthdate = ref('')
const phone = ref('')
const password = ref('')
const role = ref('')
const error = ref('')
const success = ref('')
const isLoading = ref(false)

const avatarInput = ref(null)
const avatarBase64 = ref(null)
const avatarPreview = ref(null)

// Prop changes
watch(() => props.show, (newVal) => {
  isOpen.value = newVal
  if(newVal) {
    error.value = ''
    success.value = ''
  }
})
// Internal state changes
watch(isOpen, (newVal) => {
  emit('update:show', newVal)
})

const triggerAvatarInput = () => {
  avatarInput.value.click()
}

const closeDialog = () => {
  isOpen.value = false
  error.value = ""
  emit('update:show', false)
  emit('close')
}

const handleFileChange = (event) => {
  const file = event.target.files[0]
  if (!file) return
  
  // Validate file size (e.g., 2MB max)
  if (file.size > 2 * 1024 * 1024) {
    error.value = 'File size is too large. Maximum size is 2MB.'
    return
  }
  
  // Validate file type
  if (!file.type.startsWith('image/')) {
    error.value = 'Only image files are allowed.'
    return
  }
  
  const reader = new FileReader()
  
  reader.onload = (e) => {
    avatarBase64.value = e.target.result
    avatarPreview.value = e.target.result
  }
  
  reader.onerror = () => {
    error.value = 'Error reading file'
  }
  
  reader.readAsDataURL(file)
}


const resetForm = () => {
  password.value = ''
  error.value = ''
  success.value = ''
}

const handleSubmit = async () => {
  isLoading.value = true
  error.value = ""
  if (!password.value) {
      error.value = 'Invalid Password'
      isLoading.value = false
      return
    }

    //made so that birthdate is not required everytime
  const load = {
      display_name: displayname.value,
      password: password.value,
      role: role.value,
    }
    if (phone.value) load.phone = phone.value
    if (birthdate.value) load.birth_date = birthdate.value
    if (avatarBase64.value) load.avatar = avatarBase64.value

  try {
    const response = await fetch(`api/v1/user`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(load),
    })
    

    if (!response.ok) {
      if(response.status === 401) {
        error.value = 'Invalid Password'
        return;
      }
      try{
       const resData = await response.json();
       error.value = resData.message || 'Failed to update settings'
      } catch (jsonError) {
        error.value ='Failed to update settings, try again.'
      }
      return
    }
    
    //Successful update
  success.value = 'User info updated successfully'
  setTimeout(() => window.location.reload(), 1500) //reload page

  } finally {
    isLoading.value = false
  }
  }
</script>
