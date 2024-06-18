using System;
using StudyBud.Data.Interfaces;
using StudyBud.Business.Interfaces;
using StudyBud.Models;
using UglyToad.PdfPig;
using UglyToad.PdfPig.Content;

namespace StudyBud.Business
{
	public class SyllabusBLL : ISyllabusBLL
	{
		private ISyllabusDAL _syllabusDal;

		public SyllabusBLL(ISyllabusDAL syllabusDal)
		{
			_syllabusDal = syllabusDal;
		}

		public bool ProcessSyllabus(MemoryStream syllabus, string userId)
		{
			syllabus.Position = 0;
			using (PdfDocument pdfDocument = PdfDocument.Open(syllabus))
			{
				List<string> pageData = new List<string>();

				foreach (var page in pdfDocument.GetPages())
				{

                    foreach (Word word in page.GetWords())
                    {
						pageData.Add(word.ToString());
                    }
                }

				foreach (var item in pageData)
				{
					//figure out how to deal with the words/parts of the document and
					//then create the syllabus model to store in the db
				}

			}


			return true;
		}
	}
}

