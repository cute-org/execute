<template>
    <TransitionRoot as="template" :show="show">
      <Dialog class="relative z-10" @close="closeDialog">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500/75 transition-opacity" />
        </TransitionChild>
  
        <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                <div class="sm:flex sm:items-start">
                  <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                    <DialogTitle as="h3" class="text-base font-semibold text-gray-900">That is a placeholder for settings popup dialog</DialogTitle>
                    <div class="mt-2">
                      <p class="text-sm text-gray-500">12!21(42"/.)/[[]=->..;`~`""]</p>
                    </div>
                    <button @click="handleLogout" class="px-4 py-2 rounded bg-red-600 hover:bg-red-500">Logout</button>
                  </div>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
    

</template>

<script setup>
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// properties for component listen to emit
const props = defineProps({
  show: {
    type: Boolean,
    required: true
  }
})

// emits to parent
const emit = defineEmits(['update:show'])

// Methods
const closeDialog = () => {
  emit('update:show', false)
}
const handleLogout = () => {
  emit('update:show', false) 

  setTimeout(() => {
    emit('navigate', 'logout')
  }, 300) 
}

</script>