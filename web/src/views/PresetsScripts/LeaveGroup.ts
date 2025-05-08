import { fetchTeamInfo } from "./GroupInfo";


export async function handleLeaveGroup(onSuccess= () => {}){
    try {
        const response = await fetch('api/v1/group/leave', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
              },
              credentials: 'include',
        });
        if (response.ok) {
            const result = await response.json();
            console.log(result.message);
            await fetchTeamInfo();
        }
        onSuccess()
    } catch (error) {
        console.error('Failed to fetch team users info', error);
    }
};