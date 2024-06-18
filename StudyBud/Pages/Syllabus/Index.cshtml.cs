using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;
using System.Xml.Linq;
using System.Security.Claims;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.AspNetCore.Identity;
using StudyBud.Models;
using StudyBud.Business.Interfaces;


namespace StudyBud.Views.Syllabus
{
	public class IndexModel : PageModel
    {
        private readonly ISyllabusBLL _syllabusBLL;
        private readonly IHttpContextAccessor _httpContextAccesor;

        public IndexModel(ISyllabusBLL syllabusBLL, IHttpContextAccessor httpContextAccessor)
        {
            _syllabusBLL = syllabusBLL;
            _httpContextAccesor = httpContextAccessor;
        }
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
                    var userId = _httpContextAccesor.HttpContext.User.FindFirstValue(ClaimTypes.NameIdentifier);
                    try
                    {
                        if(_syllabusBLL.ProcessSyllabus(memoryStream, userId))
                        {
                            return Page();
                        }

                    }
                    catch (Exception ex)
                    {
                        //log error and return error page
                        return Page();
                    }


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
