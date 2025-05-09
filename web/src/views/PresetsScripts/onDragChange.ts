import { el } from "date-fns/locale";
import { fetchTeamInfo } from "./GroupInfo";

export async function onDragChange(event, newStepId) {
    if (event.added) {
        const task = event.added.element;
        let action;
        
        const currentStep = task.step || newStepId;

        //todo - inprogress
        if (currentStep == 1 && newStepId == 2) {
            action = "+1"
        }

        //inprogress - completed
        else if (currentStep == 2 && newStepId == 3) {
            action = "+1"
        }

        //inprogress - todo
        else if (currentStep == 2 && newStepId == 1) {
            action = "-1"
        }

        //completed - inprogress
        else if (currentStep == 3 && newStepId == 2) {
            action = "-1"
        }


        //todo - completed
        else if (currentStep == 1 && newStepId == 3) {
            await updateTaskStep(task.id, "+1");
            await updateTaskStep(task.id, "+1");
            return;
        }
        //completed - todo
        else if (currentStep == 3 && newStepId == 1) {
            await updateTaskStep(task.id, "-1");
            await updateTaskStep(task.id, "-1");
            return;
        }
        //samecolumn
        else {
            return;
        }

        await updateTaskStep(task.id, action);
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