function findConfidenceWeight(timeMin: number) {
    const slope = 0.8;
    const midPoint = 7.5; // min
    const weight = 1 / (1 + Math.exp(-1 * slope * (timeMin - midPoint)));
    return weight;
}

function estimatePace(w: number, actualPace: number) {
    const avgPace = 1; // pages/min
    const estimate = (1 - w) * avgPace + w * actualPace; // when w = 0, fully use avg | when w = 1, fully use actual
    return estimate;
}

export function getPace(timeMin: number, actualPace: number) {
    const w = findConfidenceWeight(timeMin);
    const est = estimatePace(w, actualPace);
    return est;
}