import os
from pdfminer.high_level import extract_text
import markdownify


def pdf_to_markdown(pdf_path):
    # Extract text from PDF
    text = extract_text(pdf_path)

    # Convert text to Markdown (basic conversion)
    markdown_text = markdownify.markdownify(text, heading_style="ATX")

    return markdown_text


def process_pdf_files(file_path):
    with open(file_path, 'r') as f:
        lines = f.readlines()

    for line in lines:
        pdf_path = line.strip().strip('"')

        # Extract base filename and create new Markdown path
        base_filename = os.path.basename(pdf_path)
        new_filename = os.path.splitext(base_filename)[0].lower() + '.md'
        new_path = os.path.join('./data/source', new_filename)

        print(f'converting {base_filename}')
        # Ensure the directory exists
        os.makedirs(os.path.dirname(new_path), exist_ok=True)

        # Convert PDF to Markdown
        markdown_content = pdf_to_markdown(pdf_path)

        # Write the Markdown content to the new file
        with open(new_path, 'w') as md_file:
            md_file.write(markdown_content)

        print(f'Converted: {pdf_path} -> {new_path}')


# Example usage
process_pdf_files('./data/source/sources.txt')
