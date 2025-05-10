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
  
          <!-- Dialog Content -->
          <div v-if="!isCreating && !isJoining" class="space-y-6">
            <h2 class="text-2xl font-bold text-center text-white">Team Details</h2>
            
            <div class="space-y-4 bg-fillingInfo p-4 rounded-lg">
              <div>
                <p class="text-sm text-gray-300">Team Name</p>
                <p class="font-medium text-lg text-white">{{ teamData.name }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-300">Join Code</p>
                <p class="font-medium text-lg font-mono text-white">{{ teamData.code }}</p>
              </div>
            </div>

              <!-- Leave button -->
              <div v-if="teamData.name !== 'No group'" class="mt-4">
                <button
                  @click="leaveGroup"
                  class="w-full bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-lg transition-colors"
                >
                  Leave Group
                </button>
              </div>

            <div v-if="!teamData.name || teamData.name === 'No group'" class="grid grid-cols-2 gap-4">
              <button
                @click="isCreating = true"
                class="flex items-center justify-center gap-2 bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors"
              >
                <span class="font-bold text-lg">+</span>
                Create New
              </button>
              <button
                @click="isJoining = true"
                class="flex items-center justify-center gap-2 bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-colors"
              >
                Join Team
              </button>
            </div>
          </div>
  
          <div v-else-if="isCreating" class="space-y-6">
            <h2 class="text-2xl font-bold text-center text-white">Create New Team</h2>
            
            <form @submit.prevent="handleCreateGroup" class="space-y-4">
              <div>
                <label for="groupName" class="block text-sm font-medium text-gray-300 mb-1">
                  Team Name
                </label>
                <input
                  id="groupName"
                  type="text"
                  v-model="newGroupName"
                  class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Enter team name"
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
                  {{ isLoading ? 'Creating...' : 'Create' }}
                </button>
              </div>
            </form>
          </div>
  
          <div v-else class="space-y-6">
            <h2 class="text-2xl font-bold text-center text-white">Join a Team</h2>
            
            <form @submit.prevent="handleJoinGroup" class="space-y-4">
              <div>
                <label for="joinCode" class="block text-sm font-medium text-gray-300 mb-1">
                  Join Code
                </label>
                <input
                  id="joinCode"
                  type="text"
                  v-model="joinCode"
                  class="w-full px-3 py-2 border border-gray-600 bg-fillingInfo text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Enter join code"
                  required
                />
              </div>
              
              <p v-if="error" class="text-red-400 text-sm">{{ error }}</p>
              
              <div class="flex gap-3">
                <button
                  type="button"
                  @click="isJoining = false"
                  class="flex-1 py-2 px-4 border border-gray-600 rounded-md text-gray-300 hover:bg-gray-700 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="isLoading"
                  class="flex-1 py-2 px-4 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors disabled:opacity-50"
                >
                  {{ isLoading ? 'Joining...' : 'Join' }}
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
  import { fetchTeamInfo, teamData as rawTeamData } from '../PresetsScripts/GroupInfo.ts';
  import { computed } from 'vue';
  import { handleLeaveGroup } from '../PresetsScripts/LeaveGroup.js';
  
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
  const newGroupName = ref('');
  const joinCode = ref('');
  const error = ref('');
  const isLoading = ref(false);
 
  //  prop changes
  watch(() => props.show, (newVal) => {
    isOpen.value = newVal;
  });
  
  //  internal state changes
  watch(isOpen, (newVal) => {
    emit('update:show', newVal);
  });

  // team info changes 
  watch(() => props.show, (newVal) => {
  isOpen.value = newVal;
  if (newVal) {
    fetchTeamInfo();
  }
});

const teamData = computed(() => rawTeamData.value); //for buttons to leave/join/create visibility

const leaveGroup = () => {
  handleLeaveGroup(
    () => {
      closeDialog();
      window.location.reload();
    }
  )
}

  
  const closeDialog = () => {
    isOpen.value = false;
    isCreating.value = false;
    isJoining.value = false;
    error.value = "";
    emit('update:show', false);
  };

  const handleCreateGroup = async () => {
    try {
      const response = await fetch(`api/v1/group`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({
              name: newGroupName.value
          }),
      });
    closeDialog()
    
  } catch (error) { 
    error.value = `Connection error: ${error.message || 'Unknown error'}`;
    console.error('Creating group error:', error);
  }       
}

const handleJoinGroup = async () => {
  try {
    const response = await fetch(`api/v1/group/join`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({
        code: joinCode.value
      }),
    });
    if (response.ok) {
      await fetchTeamInfo();
        closeDialog();
    }
    closeDialog()
  } catch (error) {
    error.value = `Connection error: ${error.message || 'Unknown error'}`;
    console.error('Joining group error', error);
  }
}
  </script>