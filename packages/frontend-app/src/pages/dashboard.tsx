function Dashboard() {
       return (
              <div className="container rounded-md p-3">
                     <h2 className="font-bold text-xl mb-5">Some fields</h2>
                     <form className="" method="POST" action="http://localhost:8082/person" encType="multipart/form-data">
                            <div className="flex mb-1">
                                   <label htmlFor="" className="min-w-[100px]">Name : </label>
                                   <input type="" name="name" className="px-2 py-1 rounded border ml-2" />
                            </div>
                            
                            <div className="flex mb-1">
                                   <label htmlFor="" className="min-w-[100px]">Password : </label>
                                   <input type="text" name="password" className="px-2 py-1 rounded border ml-2" />
                            </div>
                            
                            <div className="flex mb-1">
                                   <label htmlFor="" className="min-w-[100px]">Age : </label>
                                   <input type="number" name="age" className="px-2 py-1 rounded border ml-2" />
                            </div>
                            
                            <div className="flex mb-1">
                                   <label htmlFor="" className="min-w-[100px]">File : </label>
                                   <input type="file" name="uploadfile" className="px-2 py-1 rounded border ml-2" />
                            </div>

                            <div className="flex mb-1">
                                   <label htmlFor="" className="min-w-[100px]">File : </label>
                                   <input type="file" name="uploadfile2" className="px-2 py-1 rounded border ml-2" />
                            </div>
                            <button type="submit" className="bg-teal-500 px-4 py-2 mt-3 rounded">Send</button>

                     </form>
              </div>
       )
}

export default Dashboard;