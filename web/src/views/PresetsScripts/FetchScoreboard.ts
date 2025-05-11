import { ref } from 'vue'

export const teamsData = ref([])

export const updateTeamData = (newTeamsData) => {
    if (newTeamsData && Array.isArray(newTeamsData)) {
      teamsData.value = newTeamsData
    } else {
      teamsData.value = [];
    }
  }
  
export async function fetchScoreboardInfo() {
  try {
    const response = await fetch('api/v1/scoreboard', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    
    if (response.ok) {
      const data = await response.json();
      console.log('Fetched scoreboard teams:', teamsData.value);

      if (data === null || !Array.isArray(data)) {
        teamsData.value = [];
      } else {
        teamsData.value = data;
      }
    } else {
      console.error('Error fetching scoreboard info:', response.status);
      teamsData.value = [];
    }
  } catch (error) {
    console.error('Failed to fetch scoreboard info', error);
    teamsData.value = [];
  }
}