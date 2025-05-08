import { ref } from 'vue'

export const teamsData = ref([])

export const updateTeamData = (newTeamsData) => {
    teamsData.value = newTeamsData
  }
  
export async function fetchScoreboardInfo() {
  try {
    const response = await fetch('http://localhost:8437/api/v1/scoreboard', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    
    if (response.ok) {
      const data = await response.json();
      teamsData.value = data;
      console.log('Fetched scoreboard teams:', teamsData.value)
    } else {
      console.error('Error fetching scoreboard info:', response.status);
    }
  } catch (error) {
    console.error('Failed to fetch scoreboard info', error);
  }
}