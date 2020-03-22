/* -----------------------------------
   Code by Brandio
------------------------------------*/
"use strict"
window.onload = function() {
    var canvas = document.getElementById('pagebg');
    paper.setup(canvas);
    var svg = paper.project.importSVG('<path d="M164.55,183.133c3.573,31.267,29.481,56.728,30.821,88.442,1.787,37.521-30.821,66.108-59.855,90.228-28.587,23.674-56.728,49.134-71.915,83.082-15.187,33.5-13.847,77.275,12.06,103.629,31.267,31.267,83.975,27.247,118.816,53.6,26.8,20.1,38.414,54.048,56.281,82.635,19.514,30.86,48.02,55.61,80.982,71.9,31.512,15.569,67.219,23.2,102.6,20.566,19.654-1.34,41.541-7.147,52.261-23.674,13.4-20.1,6.253-47.348-5.807-67.895s-28.587-39.754-33.5-62.981c-11.167-49.581,31.714-94.249,73.255-123.729,41.541-29.034,90.675-59.854,100.055-109.882C630.879,335.9,591.125,285.868,547.8,253.261c-25.907-19.207-54.941-37.074-69.235-66.108-11.614-23.674-10.72-51.814-20.1-76.828-12.507-34.394-45.114-60.748-81.295-66.108-66.555-9.827-78.615,44.221-126.409,67.448C216.365,128.192,157.4,118.812,164.55,183.133Z" />');

    var path = new paper.Path();
    var orgSeg = [];
    var pathOptions = {color: '#4519ff',opacity: 0.3,angle: -30, position: 'TL',x: -300,y: -500,scale: 2};

    var moveSize = 5;
    var moveSpeed = 1000;

    initializePath();

    function initializePath() {
        path.segments = [];
        path = svg;
        path.fillColor = pathOptions.color;
        path.fillColor.alpha = pathOptions.opacity;

        path.scale(pathOptions.scale);
        path.rotate(pathOptions.angle);
       
        fixSize();
        
        //path.smooth({ type: 'continuous' });

        for (var i = 0; i < path.segments.length; i++) {
            orgSeg.push([path.segments[i].point.x,path.segments[i].point.y]);
        }
    }
    function fixSize(){
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
        paper.view.bounds.width = canvas.width;
        paper.view.bounds.height = canvas.height;

        var newWidth = 1139;
        var newHeight = 1299;

        if(paper.view.bounds.width<1500){
            newWidth = (1140*paper.view.bounds.width)/1500;
            newHeight = (1300*newWidth)/1140;
            path.bounds.width = newWidth;
            path.bounds.height = newHeight;
        }

        var posX = 0;
        var posY = 0;

        posX = paper.view.bounds.topLeft.x + (path.bounds.width/2) + (1140-newWidth)/4;
        posY = paper.view.bounds.topLeft.y + (path.bounds.height/2) + (1300-newHeight)/3;
        
        path.position = new paper.Point(posX + pathOptions.x,posY + pathOptions.y);
    }
    paper.view.onFrame = function(event) {
        for (var i = 0; i < path.segments.length; i++) {
            var xPos = Math.cos(event.count*(i+1)/moveSpeed)*moveSize;
            var yPos = Math.sin(event.count*(i+1)/moveSpeed)*moveSize;
            path.segments[i].point.x = orgSeg[i][0] + xPos;
            path.segments[i].point.y = orgSeg[i][1] + yPos;
        }
    }
    paper.view.onResize = function(event) {
        fixSize();
    }
}