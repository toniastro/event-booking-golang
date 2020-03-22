/* -----------------------------------
   Code by Brandio
------------------------------------*/
"use strict"
window.onload = function() {
    var canvas = document.getElementById('pagebg');
    paper.setup(canvas);
    var svg = [paper.project.importSVG('<path d="M290.423,334.534c7.635,65.058,62.989,118.033,65.852,184.02,3.817,78.069-65.852,137.55-127.886,187.737-61.08,49.258-121.206,102.233-153.654,172.867-32.449,69.7-29.586,160.785,25.768,215.619,66.806,65.057,179.423,56.693,253.864,111.527,57.263,41.823,82.076,112.457,120.251,171.938,82.076,126.4,238.594,203.537,392.248,192.384,41.992-2.788,88.757-14.87,111.662-49.258,28.631-41.823,13.361-98.516-12.407-141.268s-61.08-82.716-71.578-131.045c-23.86-103.163,67.761-196.1,156.517-257.442,88.757-60.41,193.738-124.539,213.78-228.631,21.951-110.6-62.989-214.69-155.563-282.535-55.354-39.964-117.388-77.139-147.928-137.55-24.814-49.258-22.905-107.809-42.947-159.855-26.722-71.563-96.392-126.4-173.7-137.55C602.5,25.047,576.735,137.5,474.617,185.831,401.13,220.219,275.153,200.7,290.423,334.534Z" />'),
    paper.project.importSVG('<path d="M28.746,756.753C24.434,742.379,365.127,381.4,884.21,525.326c212.205,62.348,418.628,98.354,551.77-82.043.627-.854,27.792,34.451,31.231,33.711,29.7-6.395,56.93-8.552,47.033-8.552l39.935-14.672-107.77,393.361H243.641C-62.351,863.3,30.184,761.425,28.746,756.753Z" />'),
    paper.project.importSVG('<path d="M293.553,708.112c11.3-32.713,52.287-59.584,80.927-76.408,139.257-82.95,283.243-25.937,423.551,21.263,240.415,81.081,606.161,76.174,757.241-158.423C1562.628,483.095,1617.806,345,1605.456,345h48.608V975.889H290.4V818.167c0-30.61.526-61.22,0-92.063A53.092,53.092,0,0,1,293.553,708.112Z" />'),
    paper.project.importSVG('<path d="M927.2,695.557l536.7-.539s-139.071-11.14-261.072-68.278c-117.33-54.982-245.98-92.714-375.348-65.4-64.684,13.655-126.853,37.373-186.506,65.4-122,57.138-261.073,68.278-261.073,68.278l536.7.539h10.6Z" />')];

    var path = [new paper.Path(),new paper.Path(),new paper.Path(),new paper.Path()];

    var orgSeg = [[],[],[],[]];

    var pathOptions = [{color: '#8E1DFF',opacity: 1,angle: -50, position: 'TL',x: -400,y: -600,scale: 1},
                   {color: '#4B0096',opacity: 0.47,angle: 0, position: 'BR',x: 400,y: 280,scale: 1.3},
                   {color: '#DBB7FF',opacity: 0.17,angle: 0, position: 'BR',x: 100,y: 270,scale: 1.3},
                   {color: '#590AA8',opacity: 0.61,angle: 180, position: 'TR',x: 300,y: -55,scale: 1}];

    var moveSize = 10;
    var moveSpeed = 500;

    initializePath();

    function initializePath() {
        for(var p=0;p<path.length;p++){
            path[p].segments = [];
            path[p] = svg[p];

            path[p].fillColor = pathOptions[p].color;
            path[p].fillColor.alpha = pathOptions[p].opacity;

            path[p].scale(pathOptions[p].scale);
            path[p].rotate(pathOptions[p].angle);
        
            fixSize();
            
            //path[p].smooth({ type: 'continuous' });

            for (var i = 0; i < path[p].segments.length; i++) {
                orgSeg[p].push([path[p].segments[i].point.x,path[p].segments[i].point.y]);
            }
        }
    }
    function fixSize(){
       canvas.width = window.innerWidth;
       canvas.height = window.innerHeight;
       paper.view.bounds.width = canvas.width;
       paper.view.bounds.height = canvas.height;

       for(var p=0;p<path.length;p++){
           var posX = 0;
           var posY = 0;
           if(pathOptions[p].position=="TL"){
               posX = paper.view.bounds.topLeft.x + path[p].bounds.width/2;
               posY = paper.view.bounds.topLeft.y + path[p].bounds.height/2;
           }else if(pathOptions[p].position=="BL"){
               posX = paper.view.bounds.bottomLeft.x + path[p].bounds.width/2;
               posY = paper.view.bounds.bottomLeft.y - path[p].bounds.height/2;
           }else if(pathOptions[p].position=="TR"){
               posX = paper.view.bounds.topRight.x - path[p].bounds.width/2;
               posY = paper.view.bounds.topRight.y + path[p].bounds.height/2;
           }else if(pathOptions[p].position=="BR"){
               posX = paper.view.bounds.bottomRight.x - path[p].bounds.width/2;
               posY = paper.view.bounds.bottomRight.y - path[p].bounds.height/2;
           }
           path[p].position = new paper.Point(posX + pathOptions[p].x,posY + pathOptions[p].y);
       }
    }
    paper.view.onFrame = function(event) {
        for(var p=0;p<path.length;p++){
            for (var i = 0; i < path[p].segments.length; i++) {
                var xPos = Math.cos(event.count*(i+1)/moveSpeed)*moveSize;
                var yPos = Math.sin(event.count*(i+1)/moveSpeed)*moveSize;
                path[p].segments[i].point.x = orgSeg[p][i][0] + xPos;
                path[p].segments[i].point.y = orgSeg[p][i][1] + yPos;
            }
        }
    }
    paper.view.onResize = function(event) {
        fixSize();
    }
}