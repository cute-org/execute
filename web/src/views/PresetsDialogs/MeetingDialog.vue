<!-- TeamDialog.vue -->
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
  
          <div class="space-y-6">
            <h2 class="text-2xl font-bold text-center text-white">Schedule a meeting</h2>
            
            <form class="space-y-4" @submit.prevent="handleMeetingSchedule">
              <div>
                <label for="meetingTime" class="block text-sm font-medium text-gray-300 mb-1">
                  Meeting time
                </label>
                <input
                  id="meetingTime"
                  type="datetime-local"
                  v-model="newMeetingTime"
                  class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Enter meeting date"
                  required
                />
              </div>
              
              <p v-if="error" class="text-red-400 text-sm">{{ error }}</p>
              
              <div class="flex gap-3">
                <button
                  type="button"
                  @click="isCreating = false"
                  class="flex-1 py-2 px-4 border border-gray-600 rounded-md text-gray-300 hover:bg-gray-700 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="isLoading"
                  class="flex-1 py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50"
                >
                  {{ isLoading ? 'Scheduling...' : 'Schedule' }}
                </button>
              </div>
            </form>
          </div>
        </DialogPanel>
      </div>
    </Dialog>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue';
  import { Dialog, DialogPanel } from '@headlessui/vue';

  
  const props = defineProps({
    show: {
      type: Boolean,
      default: false
    }
  });
  
  const emit = defineEmits(['update:show']);
  
  const isOpen = ref(props.show);
  const isCreating = ref(false);
  const isJoining = ref(false);
  const error = ref('');
  const isLoading = ref(false);
  const newMeetingTime = ref('')
 
  
  const API_BASE_URL = 'http://localhost:8437/api/v1';
  //  prop changes
  watch(() => props.show, (newVal) => {
    isOpen.value = newVal;
  });
  
  //  internal state changes
  watch(isOpen, (newVal) => {
    emit('update:show', newVal);
  });

  
  const closeDialog = () => {
    isOpen.value = false;
    isCreating.value = false;
    isJoining.value = false;
    error.value = "";
    emit('update:show', false);
  };

  const handleMeetingSchedule = async () => {
    try {
        const formattedTime = new Date(newMeetingTime.value).toISOString();
        const response = await fetch(`${API_BASE_URL}/group/meeting`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
                time: formattedTime
            })
        });
    
    if (!response.ok) {
        if(response.status === 403) {
            error.value = 'Only group creator can schedule a meeting'
            return;
        }
        const errorData = await response.json().catch(() => ({}));
        error.value = errorData.message || 'Failed to schedule meeting';
        return
    } 
    
    closeDialog();
    window.location.reload();
    } catch (error) {
        error.value = `Connection error: ${error.message || 'Unknown error'}`;
        console.error('Meeting schedule error', error);
    }
    
  } 
  </script>