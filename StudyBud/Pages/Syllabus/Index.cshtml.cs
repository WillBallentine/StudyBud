using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;
using System.Xml.Linq;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using StudyBud.Models;


namespace StudyBud.Views.Syllabus
{
	public class IndexModel : PageModel
    {
        [BindProperty]
        public FileUpload FileUpload { get; set; }

        public void OnGet()
        {
        }

        public async Task<IActionResult> OnPostUploadAsync()
        {
            using (var memoryStream = new MemoryStream())
            {
                
                await FileUpload.FormFile.CopyToAsync(memoryStream);

                // Process/Upload the file if less than 2 MB
                if (memoryStream.Length < 2097152)
                {
                    var file = new Models.Syllabus()
                    {
                        Content = memoryStream.ToArray()
                    };

                    //process syllabus
                    //update model
                    //call DAL to save to db

                }
                else
                {
                    ModelState.AddModelError("File", "The file is too large.");
                }
            }

            return Page();
        }


    }


    public class FileUpload
    {
        [Required]
        [Display(Name = "File")]
        public IFormFile FormFile { get; set; }
    }
}
