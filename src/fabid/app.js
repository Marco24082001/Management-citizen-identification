const express = require('express');
const cors = require('cors');
const app = express();
app.use(cors());
const PORT = 3000;
app.use(cors({
  origin: '*',
  methods: ['GET','POST','DELETE','UPDATE','PUT','PATCH']
}))
  

app.use(express.json({limit: '50mb'}));
app.use(express.urlencoded({limit: '50mb', extended: true, parameterLimit: 50000}));

const identicationRouter = require('./routes/identication.route');
const adminRouter = require('./routes/admin.route');
app.use('/identications', identicationRouter);
app.use('/admins', adminRouter);

app.listen(PORT, (error) =>{
    if(!error)
        console.log("Server is Successfully Running, and App is listening on port "+ PORT);
    else 
        console.log("Error occurred, server can't start", error);
    }
);


