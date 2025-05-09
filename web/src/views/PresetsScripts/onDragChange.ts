import { fetchTeamInfo } from "./GroupInfo";


export async function onDragChange(event, newStepId) {
    //if no tasks - exit
    if (!event.added) return;
    //what task was dropped
    const task = event.added.element;
    //Current task step
    const currentStep = task.step;
    //If task is not moved - exit 
    if (currentStep === newStepId) return;
    //Calculate step
    const stepChange = newStepId - currentStep;
    //Forward / Backwards
    const action = stepChange > 0 ? "+1" : "-1";

    //Loop number of steps to move
    for (let i = 0; i < Math.abs(stepChange); i++) {
        const success = await updateTaskStep(task.id, action); //Move task call updateTaskStep
        if (!success) break; //If failed, stop 
        task.step += action === "+1" ? 1 : -1; //Update local after each update
    }
}



export async function updateTaskStep(taskId, action) {
    try {
        const response = await fetch('http://localhost:8437/api/v1/task', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
                taskId: taskId,
                action: action
            })
        });
        if (response.ok) {
            await fetchTeamInfo();
            return true;
        } else {
            console.error('Failed to update task:', response.status)
            await fetchTasks();
            return false;
        } 
    
    }catch (error) {
        console.error('Error updating task step', error);   
        await fetchTasks();
        return false;   
    }
}